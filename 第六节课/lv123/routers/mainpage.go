package routers

import (
	"log"
	"net/http"
	"project6/tools"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func mainPageGet(ctx *gin.Context) {
	m := map[int](*[](*tools.CommentForRender)){}

	// to pass for render
	var renderComments [](*tools.CommentForRender)

	// to store all user comments info, for satisfication
	var allUserComments []tools.Comment

	rows, err := tools.GetDBHandler().Query(`SELECT comment.id AS comment_id, user.id AS user_id, comment.message AS message, comment.parent AS parent, comment.created_at AS created_at,
	 user.username AS username FROM comment, user WHERE comment.user_id=user.id ORDER BY comment.id ASC`)
	if err != nil {
		log.Panicf("failed to Query: %v\n", err)
	}
	for rows.Next() {
		var tmpUserComment tools.Comment
		if err := rows.Scan(&tmpUserComment.CommentID, &tmpUserComment.UserID, &tmpUserComment.Comment, &tmpUserComment.Parent, &tmpUserComment.CreatedAt, &tmpUserComment.Owner); err != nil {
			log.Panicf("failed to Scan: %v\n", err)
		}
		allUserComments = append(allUserComments, tmpUserComment)
	}

	// fill renderComments
	for _, v := range allUserComments {
		// if it is a top comment, put it in renderComments
		if v.Parent == nil {
			tmpRenderComment := tools.CommentForRender{
				CommentID:   v.CommentID,
				Owner:       v.Owner,
				Content:     v.Comment,
				CreatedAt:   v.CreatedAt,
				SonComments: make([](*tools.CommentForRender), 0),
			}
			renderComments = append(renderComments, &tmpRenderComment)
			m[v.CommentID] = &tmpRenderComment.SonComments
		} else {
			tmpRenderComment := tools.CommentForRender{
				Owner:       v.Owner,
				Content:     v.Comment,
				CreatedAt:   v.CreatedAt,
				SonComments: nil,
			}
			(*m[*v.Parent]) = append((*m[*v.Parent]), &tmpRenderComment)
		}
	}

	user_, _ := ctx.Get("user")
	user := user_.(*tools.User)

	ctx.HTML(http.StatusOK, "mainpage.tmpl", gin.H{
		"topcomments": renderComments,
		"user":        user,
	})
}

func mainPagePost(ctx *gin.Context) {
	postType, ok := ctx.GetPostForm("post_type")
	if !ok {
		ctx.String(http.StatusBadRequest, "invalid post data")
		return
	}

	switch postType {
	case "new_comment":
		comment, ok := ctx.GetPostForm("new_comment")
		if !ok || !tools.CheckNewCommentValid(comment) {
			ctx.String(http.StatusBadRequest, "invalid post data")
			return
		}
		doInsertMainPagePostForTopComment(ctx, comment)
	case "append_comment":
		appendIDStr, ok1 := ctx.GetPostForm("append_id")
		appendComment, ok2 := ctx.GetPostForm("append_comment")
		appendID, err := strconv.Atoi(appendIDStr)
		if !ok1 || !ok2 || err != nil || !tools.CheckAppendCommentValid(appendComment) {
			ctx.String(http.StatusBadRequest, "invalid post data")
			return
		}
		doInsertMainPagePostForAppendComment(ctx, appendID, appendComment)
	default:
		ctx.String(http.StatusBadRequest, "invalid post data")
		return
	}
}

func doInsertMainPagePostForTopComment(ctx *gin.Context, content string) {
	user_, _ := ctx.Get("user")
	user := user_.(*tools.User)
	_, err := tools.GetDBHandler().Exec(`INSERT INTO comment(user_id, message, created_at, parent)
	VALUES(?, ?, ?, ?)`, user.ID, content, time.Now().Format("2006-01-02 15:04:05"), nil)
	if err != nil {
		log.Fatalf("failed to INSERT: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "new_comment.success.tmpl", nil)
}

func doInsertMainPagePostForAppendComment(ctx *gin.Context, num int, msg string) {
	user_, _ := ctx.Get("user")
	user := user_.(*tools.User)
	// check if the comment(id = num) is a top comment
	var parent *int
	row := tools.GetDBHandler().QueryRow("SELECT parent FROM comment WHERE id=?", num)
	if row.Err() != nil {
		log.Panicf("failed to SELECT: %v\n", row.Err())
	}
	if err := row.Scan(&parent); err != nil {
		log.Panicf("failed to Scan: %v\n", err)
	}

	if parent != nil {
		ctx.String(http.StatusBadRequest, "invalid post data")
	}

	_, err := tools.GetDBHandler().Exec(`INSERT INTO comment(user_id, message, created_at, parent)
	VALUES(?, ?, ?, ?)`, user.ID, msg, time.Now().Format("2006-01-02 15:04:05"), num)
	if err != nil {
		log.Panicf("failed to INSERT: %v\n", err)
	}

	ctx.HTML(http.StatusOK, "new_comment.success.tmpl", nil)
}
