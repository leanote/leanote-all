// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/leanote/leanote/app"
	controllers "github.com/leanote/leanote/app/controllers"
	admin "github.com/leanote/leanote/app/controllers/admin"
	api "github.com/leanote/leanote/app/controllers/api"
	member "github.com/leanote/leanote/app/controllers/member"
	info "github.com/leanote/leanote/app/info"
	_ "github.com/leanote/leanote/app/lea/binder"
	_ "github.com/leanote/leanote/app/lea/captcha"
	_ "github.com/leanote/leanote/app/lea/i18n"
	_ "github.com/leanote/leanote/app/service"
	_ "github.com/leanote/leanote/app/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.BaseController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "E404",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RenderRe",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "re", Type: reflect.TypeOf((*info.Re)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Note)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "online", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListNotes",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListTrashNotes",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNoteAndContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNoteContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateNoteOrContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteOrContent", Type: reflect.TypeOf((*info.NoteOrContent)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "isShared", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTrash",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "MoveNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CopyNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CopySharedNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fromUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SearchNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "key", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SearchNoteByTags",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tags", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ToPdf",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "appKey", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ExportPdf",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetNote2Blog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "isBlog", Type: reflect.TypeOf((*bool)(nil)) },
					&revel.MethodArg{Name: "isTop", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Index)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Default",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Suggestion",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "addr", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "suggestion", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Blog)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "E",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Tags",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Tag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Archives",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "cateId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "month", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Cate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Post",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Single",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetPostStat",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetLikes",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetLikesAndComments",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncReadNum",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "LikePost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetComments",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteComment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "commentId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CommentPost",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "content", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "toCommentId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "LikeComment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "commentId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListCateLatest",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "callback", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Captcha)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Get",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					43: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Oauth)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GithubCallback",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "code", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Account",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tab", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateUsername",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdatePwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "oldPwd", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "theme", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SendRegisterEmail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "content", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "toEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ReSendActiveEmail",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateEmail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ActiveEmail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateColumnWidth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookWidth", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "noteListWidth", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "mdEditorWidth", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateLeftIsMin",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "leftIsMin", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Attach)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "UploadAttach",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteAttach",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "attachId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetAttachs",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Download",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "attachId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DownloadAll",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Notebook)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebook", Type: reflect.TypeOf((*info.Notebook)(nil)) },
					&revel.MethodArg{Name: "i", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNotebooks",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "parentNotebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateNotebookTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DragNotebooks",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "data", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetNotebook2Blog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "isBlog", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Tag)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "UpdateTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Album)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetAlbums",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteAlbum",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "albumId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddAlbum",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateAlbum",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "albumId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Auth)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "from", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoLogin",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "captcha", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Demo",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "from", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "iu", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoRegister",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "iu", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FindPassword",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoFindPassword",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FindPassword2",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FindPasswordUpdate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Share)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "AddShareNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emails", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddShareNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emails", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListShareNotes",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "userId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetShareNoteContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "sharedUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListNoteShareUserInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListNotebookShareUserInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateShareNotePerm",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "toUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateShareNotebookPerm",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "toUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "toUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "toUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNoteBySharedUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fromUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNotebookBySharedUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fromUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteUserShareNoteAndNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fromUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddShareNoteGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNoteGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateShareNoteGroupPerm",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddShareNotebookGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteShareNotebookGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateShareNotebookGroupPerm",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "perm", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.File)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "UploadBlogLogo",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PasteImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UploadAvatar",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UploadImageLeaui",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "albumId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetImages",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "albumId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "key", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "page", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateImageTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "OutputImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CopyImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "toUserId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CopyHttpImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "src", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.NoteContentHistory)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ListHistories",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminBaseController)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*api.ApiBaseContrller)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*member.MemberBaseController)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers.Preview)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Tag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Tags",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Archives",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "month", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Cate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Post",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Single",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userIdOrEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminData)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Backup",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Restore",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "createdTime", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "createdTime", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateRemark",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "createdTime", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "remark", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Download",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "createdTime", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminEmail)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Email",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Blog",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoBlogTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "recommendTags", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "newTags", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Demo",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoDemo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "demoUsername", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "demoPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ToImage",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoToImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "toImageBinPath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Set",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "emailHost", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailPort", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailUsername", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Template",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SendEmailToEmails",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "sendEmails", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailSubject", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailBody", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verified", Type: reflect.TypeOf((*bool)(nil)) },
					&revel.MethodArg{Name: "saveAsOldEmail", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SendToUsers2",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "emails", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailSubject", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailBody", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verified", Type: reflect.TypeOf((*bool)(nil)) },
					&revel.MethodArg{Name: "saveAsOldEmail", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SendEmailDialog",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "emails", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SendToUsers",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userFilterEmail", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "userFilterWhiteList", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "userFilterBlackList", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailSubject", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "latestEmailBody", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verified", Type: reflect.TypeOf((*bool)(nil)) },
					&revel.MethodArg{Name: "saveAsOldEmail", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteEmails",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "ids", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "sorter", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminUser)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "sorter", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ResetPwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoResetPwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminBlog)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "sorter", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetRecommend",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "recommend", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "T",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "t", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetView",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "view", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminSetting)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Email",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Blog",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoBlogTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "recommendTags", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "newTags", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ShareNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "registerSharedUserId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "registerSharedNotebookPerms", Type: reflect.TypeOf((*[]int)(nil)) },
					&revel.MethodArg{Name: "registerSharedNotePerms", Type: reflect.TypeOf((*[]int)(nil)) },
					&revel.MethodArg{Name: "registerSharedNotebookIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "registerSharedNoteIds", Type: reflect.TypeOf((*[]string)(nil)) },
					&revel.MethodArg{Name: "registerCopyNoteIds", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Demo",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoDemo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "demoUsername", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "demoPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ExportPdf",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "path", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoSiteUrl",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "siteUrl", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SubDomain",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoSubDomain",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteSubDomain", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "blogSubDomain", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "leaSubDomain", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "blackSubDomains", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "allowCustomDomain", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "blackCustomDomains", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "OpenRegister",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "openRegister", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "HomePage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "homePage", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Mongodb",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "mongodumpPath", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "mongorestorePath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UploadSize",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "uploadImageSize", Type: reflect.TypeOf((*float64)(nil)) },
					&revel.MethodArg{Name: "uploadAvatarSize", Type: reflect.TypeOf((*float64)(nil)) },
					&revel.MethodArg{Name: "uploadBlogLogoSize", Type: reflect.TypeOf((*float64)(nil)) },
					&revel.MethodArg{Name: "uploadAttachSize", Type: reflect.TypeOf((*float64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*admin.AdminUpgrade)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "UpgradeBlog",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpgradeBetaToBeta2",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpgradeBeta3ToBeta4",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiTag)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetSyncTags",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "afterUsn", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "maxEntry", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "usn", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiUser)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Info",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateUsername",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdatePwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "oldPwd", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetSyncState",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateLogo",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiAuth)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pwd", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiFile)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetAttach",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "fileId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetAllAttachs",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiNotebook)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetSyncNotebooks",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "afterUsn", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "maxEntry", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNotebooks",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "parentNotebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "seq", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "parentNotebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "seq", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "usn", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteNotebook",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "usn", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*api.ApiNote)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetSyncNotes",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "afterUsn", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "maxEntry", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNotes",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "notebookId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetTrashNotes",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNoteAndContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetNoteContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteOrContent", Type: reflect.TypeOf((*info.ApiNote)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateNote",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteOrContent", Type: reflect.TypeOf((*info.ApiNote)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTrash",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "usn", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ExportPdf",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*member.MemberIndex)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "T",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "t", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetView",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "view", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*member.MemberUser)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Username",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Email",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Password",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Avatar",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*member.MemberBlog)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "sorter", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "keywords", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateBlogUrlTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "urlTitle", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateBlogAbstract",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoUpdateBlogAbstract",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "noteId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "imgSrc", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "desc", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "abstract", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Base",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Comment",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Paging",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Cate",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpateCateIds",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "cateIds", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateCateUrlTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "cateId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "urlTitle", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoAddOrUpdateSingle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "content", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddOrUpdateSingle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SortSingles",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "singleIds", Type: reflect.TypeOf((*[]string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteSingle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateSingleUrlTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "singleId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "urlTitle", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Single",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Theme",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "isNew", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetTplContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filename", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateTplContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filename", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "content", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTpl",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filename", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListThemeImages",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteThemeImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filename", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UploadThemeImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ActiveTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PublicTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ExportTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ImportTheme",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "InstallTheme",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "themeId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "NewTheme",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetUserBlogBase",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userBlog", Type: reflect.TypeOf((*info.UserBlogBase)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetUserBlogComment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userBlog", Type: reflect.TypeOf((*info.UserBlogComment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetUserBlogStyle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "userBlog", Type: reflect.TypeOf((*info.UserBlogStyle)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetUserBlogPaging",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "perPageSize", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "sortField", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "isAsc", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*member.MemberGroup)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateGroupTitle",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteGroup",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "groupId", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "userId", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
