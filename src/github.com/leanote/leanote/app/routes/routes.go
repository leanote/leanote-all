// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBaseController struct {}
var BaseController tBaseController


func (_ tBaseController) E404(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.E404", args).URL
}

func (_ tBaseController) RenderRe(
		re interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "re", re)
	return revel.MainRouter.Reverse("BaseController.RenderRe", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tAlbum struct {}
var Album tAlbum


func (_ tAlbum) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Album.Index", args).URL
}

func (_ tAlbum) GetAlbums(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Album.GetAlbums", args).URL
}

func (_ tAlbum) DeleteAlbum(
		albumId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "albumId", albumId)
	return revel.MainRouter.Reverse("Album.DeleteAlbum", args).URL
}

func (_ tAlbum) AddAlbum(
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Album.AddAlbum", args).URL
}

func (_ tAlbum) UpdateAlbum(
		albumId string,
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "albumId", albumId)
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Album.UpdateAlbum", args).URL
}


type tAttach struct {}
var Attach tAttach


func (_ tAttach) UploadAttach(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Attach.UploadAttach", args).URL
}

func (_ tAttach) DeleteAttach(
		attachId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "attachId", attachId)
	return revel.MainRouter.Reverse("Attach.DeleteAttach", args).URL
}

func (_ tAttach) GetAttachs(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Attach.GetAttachs", args).URL
}

func (_ tAttach) Download(
		attachId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "attachId", attachId)
	return revel.MainRouter.Reverse("Attach.Download", args).URL
}

func (_ tAttach) DownloadAll(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Attach.DownloadAll", args).URL
}


type tShare struct {}
var Share tShare


func (_ tShare) AddShareNote(
		noteId string,
		emails []string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "emails", emails)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.AddShareNote", args).URL
}

func (_ tShare) AddShareNotebook(
		notebookId string,
		emails []string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "emails", emails)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.AddShareNotebook", args).URL
}

func (_ tShare) ListShareNotes(
		notebookId string,
		userId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("Share.ListShareNotes", args).URL
}

func (_ tShare) GetShareNoteContent(
		noteId string,
		sharedUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "sharedUserId", sharedUserId)
	return revel.MainRouter.Reverse("Share.GetShareNoteContent", args).URL
}

func (_ tShare) ListNoteShareUserInfo(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Share.ListNoteShareUserInfo", args).URL
}

func (_ tShare) ListNotebookShareUserInfo(
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Share.ListNotebookShareUserInfo", args).URL
}

func (_ tShare) UpdateShareNotePerm(
		noteId string,
		perm int,
		toUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "perm", perm)
	revel.Unbind(args, "toUserId", toUserId)
	return revel.MainRouter.Reverse("Share.UpdateShareNotePerm", args).URL
}

func (_ tShare) UpdateShareNotebookPerm(
		notebookId string,
		perm int,
		toUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "perm", perm)
	revel.Unbind(args, "toUserId", toUserId)
	return revel.MainRouter.Reverse("Share.UpdateShareNotebookPerm", args).URL
}

func (_ tShare) DeleteShareNote(
		noteId string,
		toUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "toUserId", toUserId)
	return revel.MainRouter.Reverse("Share.DeleteShareNote", args).URL
}

func (_ tShare) DeleteShareNotebook(
		notebookId string,
		toUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "toUserId", toUserId)
	return revel.MainRouter.Reverse("Share.DeleteShareNotebook", args).URL
}

func (_ tShare) DeleteShareNoteBySharedUser(
		noteId string,
		fromUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "fromUserId", fromUserId)
	return revel.MainRouter.Reverse("Share.DeleteShareNoteBySharedUser", args).URL
}

func (_ tShare) DeleteShareNotebookBySharedUser(
		notebookId string,
		fromUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "fromUserId", fromUserId)
	return revel.MainRouter.Reverse("Share.DeleteShareNotebookBySharedUser", args).URL
}

func (_ tShare) DeleteUserShareNoteAndNotebook(
		fromUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fromUserId", fromUserId)
	return revel.MainRouter.Reverse("Share.DeleteUserShareNoteAndNotebook", args).URL
}

func (_ tShare) AddShareNoteGroup(
		noteId string,
		groupId string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.AddShareNoteGroup", args).URL
}

func (_ tShare) DeleteShareNoteGroup(
		noteId string,
		groupId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("Share.DeleteShareNoteGroup", args).URL
}

func (_ tShare) UpdateShareNoteGroupPerm(
		noteId string,
		groupId string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.UpdateShareNoteGroupPerm", args).URL
}

func (_ tShare) AddShareNotebookGroup(
		notebookId string,
		groupId string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.AddShareNotebookGroup", args).URL
}

func (_ tShare) DeleteShareNotebookGroup(
		notebookId string,
		groupId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("Share.DeleteShareNotebookGroup", args).URL
}

func (_ tShare) UpdateShareNotebookGroupPerm(
		notebookId string,
		groupId string,
		perm int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "perm", perm)
	return revel.MainRouter.Reverse("Share.UpdateShareNotebookGroupPerm", args).URL
}


type tTag struct {}
var Tag tTag


func (_ tTag) UpdateTag(
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Tag.UpdateTag", args).URL
}

func (_ tTag) DeleteTag(
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Tag.DeleteTag", args).URL
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Get(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Get", args).URL
}


type tNote struct {}
var Note tNote


func (_ tNote) Index(
		noteId string,
		online string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "online", online)
	return revel.MainRouter.Reverse("Note.Index", args).URL
}

func (_ tNote) ListNotes(
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Note.ListNotes", args).URL
}

func (_ tNote) ListTrashNotes(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Note.ListTrashNotes", args).URL
}

func (_ tNote) GetNoteAndContent(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Note.GetNoteAndContent", args).URL
}

func (_ tNote) GetNoteAndContentBySrc(
		src string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "src", src)
	return revel.MainRouter.Reverse("Note.GetNoteAndContentBySrc", args).URL
}

func (_ tNote) GetNoteContent(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Note.GetNoteContent", args).URL
}

func (_ tNote) UpdateNoteOrContent(
		noteOrContent interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteOrContent", noteOrContent)
	return revel.MainRouter.Reverse("Note.UpdateNoteOrContent", args).URL
}

func (_ tNote) DeleteNote(
		noteIds []string,
		isShared bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteIds", noteIds)
	revel.Unbind(args, "isShared", isShared)
	return revel.MainRouter.Reverse("Note.DeleteNote", args).URL
}

func (_ tNote) DeleteTrash(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Note.DeleteTrash", args).URL
}

func (_ tNote) MoveNote(
		noteIds []string,
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteIds", noteIds)
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Note.MoveNote", args).URL
}

func (_ tNote) CopyNote(
		noteIds []string,
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteIds", noteIds)
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Note.CopyNote", args).URL
}

func (_ tNote) CopySharedNote(
		noteIds []string,
		notebookId string,
		fromUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteIds", noteIds)
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "fromUserId", fromUserId)
	return revel.MainRouter.Reverse("Note.CopySharedNote", args).URL
}

func (_ tNote) SearchNote(
		key string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "key", key)
	return revel.MainRouter.Reverse("Note.SearchNote", args).URL
}

func (_ tNote) SearchNoteByTags(
		tags []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tags", tags)
	return revel.MainRouter.Reverse("Note.SearchNoteByTags", args).URL
}

func (_ tNote) ToPdf(
		noteId string,
		appKey string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "appKey", appKey)
	return revel.MainRouter.Reverse("Note.ToPdf", args).URL
}

func (_ tNote) ExportPdf(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Note.ExportPdf", args).URL
}

func (_ tNote) SetNote2Blog(
		noteIds []string,
		isBlog bool,
		isTop bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteIds", noteIds)
	revel.Unbind(args, "isBlog", isBlog)
	revel.Unbind(args, "isTop", isTop)
	return revel.MainRouter.Reverse("Note.SetNote2Blog", args).URL
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Login(
		email string,
		from string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "from", from)
	return revel.MainRouter.Reverse("Auth.Login", args).URL
}

func (_ tAuth) DoLogin(
		email string,
		pwd string,
		captcha string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pwd", pwd)
	revel.Unbind(args, "captcha", captcha)
	return revel.MainRouter.Reverse("Auth.DoLogin", args).URL
}

func (_ tAuth) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Logout", args).URL
}

func (_ tAuth) Demo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Demo", args).URL
}

func (_ tAuth) Register(
		from string,
		iu string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "from", from)
	revel.Unbind(args, "iu", iu)
	return revel.MainRouter.Reverse("Auth.Register", args).URL
}

func (_ tAuth) DoRegister(
		email string,
		pwd string,
		iu string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pwd", pwd)
	revel.Unbind(args, "iu", iu)
	return revel.MainRouter.Reverse("Auth.DoRegister", args).URL
}

func (_ tAuth) FindPassword(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.FindPassword", args).URL
}

func (_ tAuth) DoFindPassword(
		email string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	return revel.MainRouter.Reverse("Auth.DoFindPassword", args).URL
}

func (_ tAuth) FindPassword2(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("Auth.FindPassword2", args).URL
}

func (_ tAuth) FindPasswordUpdate(
		token string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("Auth.FindPasswordUpdate", args).URL
}


type tFile struct {}
var File tFile


func (_ tFile) UploadBlogLogo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("File.UploadBlogLogo", args).URL
}

func (_ tFile) PasteImage(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("File.PasteImage", args).URL
}

func (_ tFile) UploadAvatar(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("File.UploadAvatar", args).URL
}

func (_ tFile) UploadImageLeaui(
		albumId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "albumId", albumId)
	return revel.MainRouter.Reverse("File.UploadImageLeaui", args).URL
}

func (_ tFile) GetImages(
		albumId string,
		key string,
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "albumId", albumId)
	revel.Unbind(args, "key", key)
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("File.GetImages", args).URL
}

func (_ tFile) UpdateImageTitle(
		fileId string,
		title string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fileId", fileId)
	revel.Unbind(args, "title", title)
	return revel.MainRouter.Reverse("File.UpdateImageTitle", args).URL
}

func (_ tFile) DeleteImage(
		fileId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fileId", fileId)
	return revel.MainRouter.Reverse("File.DeleteImage", args).URL
}

func (_ tFile) OutputImage(
		noteId string,
		fileId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "fileId", fileId)
	return revel.MainRouter.Reverse("File.OutputImage", args).URL
}

func (_ tFile) CopyImage(
		userId string,
		fileId string,
		toUserId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	revel.Unbind(args, "fileId", fileId)
	revel.Unbind(args, "toUserId", toUserId)
	return revel.MainRouter.Reverse("File.CopyImage", args).URL
}

func (_ tFile) CopyHttpImage(
		src string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "src", src)
	return revel.MainRouter.Reverse("File.CopyHttpImage", args).URL
}


type tIndex struct {}
var Index tIndex


func (_ tIndex) Default(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Index.Default", args).URL
}

func (_ tIndex) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Index.Index", args).URL
}

func (_ tIndex) Suggestion(
		addr string,
		suggestion string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "addr", addr)
	revel.Unbind(args, "suggestion", suggestion)
	return revel.MainRouter.Reverse("Index.Suggestion", args).URL
}


type tBlog struct {}
var Blog tBlog


func (_ tBlog) E(
		userIdOrEmail string,
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Blog.E", args).URL
}

func (_ tBlog) Tags(
		userIdOrEmail string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	return revel.MainRouter.Reverse("Blog.Tags", args).URL
}

func (_ tBlog) Tag(
		userIdOrEmail string,
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Blog.Tag", args).URL
}

func (_ tBlog) Archives(
		userIdOrEmail string,
		cateId string,
		year int,
		month int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "cateId", cateId)
	revel.Unbind(args, "year", year)
	revel.Unbind(args, "month", month)
	return revel.MainRouter.Reverse("Blog.Archives", args).URL
}

func (_ tBlog) Cate(
		userIdOrEmail string,
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Blog.Cate", args).URL
}

func (_ tBlog) Index(
		userIdOrEmail string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	return revel.MainRouter.Reverse("Blog.Index", args).URL
}

func (_ tBlog) Post(
		userIdOrEmail string,
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Blog.Post", args).URL
}

func (_ tBlog) Single(
		userIdOrEmail string,
		singleId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "singleId", singleId)
	return revel.MainRouter.Reverse("Blog.Single", args).URL
}

func (_ tBlog) Search(
		userIdOrEmail string,
		keywords string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "keywords", keywords)
	return revel.MainRouter.Reverse("Blog.Search", args).URL
}

func (_ tBlog) GetPostStat(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Blog.GetPostStat", args).URL
}

func (_ tBlog) GetLikes(
		noteId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.GetLikes", args).URL
}

func (_ tBlog) GetLikesAndComments(
		noteId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.GetLikesAndComments", args).URL
}

func (_ tBlog) IncReadNum(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Blog.IncReadNum", args).URL
}

func (_ tBlog) LikePost(
		noteId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.LikePost", args).URL
}

func (_ tBlog) GetComments(
		noteId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.GetComments", args).URL
}

func (_ tBlog) DeleteComment(
		noteId string,
		commentId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "commentId", commentId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.DeleteComment", args).URL
}

func (_ tBlog) CommentPost(
		noteId string,
		content string,
		toCommentId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "content", content)
	revel.Unbind(args, "toCommentId", toCommentId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.CommentPost", args).URL
}

func (_ tBlog) LikeComment(
		commentId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "commentId", commentId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.LikeComment", args).URL
}

func (_ tBlog) ListCateLatest(
		notebookId string,
		callback string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "callback", callback)
	return revel.MainRouter.Reverse("Blog.ListCateLatest", args).URL
}


type tNotebook struct {}
var Notebook tNotebook


func (_ tNotebook) Index(
		notebook interface{},
		i int,
		name string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebook", notebook)
	revel.Unbind(args, "i", i)
	revel.Unbind(args, "name", name)
	return revel.MainRouter.Reverse("Notebook.Index", args).URL
}

func (_ tNotebook) GetNotebooks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Notebook.GetNotebooks", args).URL
}

func (_ tNotebook) DeleteNotebook(
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Notebook.DeleteNotebook", args).URL
}

func (_ tNotebook) AddNotebook(
		notebookId string,
		title string,
		parentNotebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "parentNotebookId", parentNotebookId)
	return revel.MainRouter.Reverse("Notebook.AddNotebook", args).URL
}

func (_ tNotebook) UpdateNotebookTitle(
		notebookId string,
		title string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "title", title)
	return revel.MainRouter.Reverse("Notebook.UpdateNotebookTitle", args).URL
}

func (_ tNotebook) DragNotebooks(
		data string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "data", data)
	return revel.MainRouter.Reverse("Notebook.DragNotebooks", args).URL
}

func (_ tNotebook) SetNotebook2Blog(
		notebookId string,
		isBlog bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "isBlog", isBlog)
	return revel.MainRouter.Reverse("Notebook.SetNotebook2Blog", args).URL
}


type tNoteContentHistory struct {}
var NoteContentHistory tNoteContentHistory


func (_ tNoteContentHistory) ListHistories(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("NoteContentHistory.ListHistories", args).URL
}


type tUser struct {}
var User tUser


func (_ tUser) Account(
		tab int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tab", tab)
	return revel.MainRouter.Reverse("User.Account", args).URL
}

func (_ tUser) UpdateUsername(
		username string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	return revel.MainRouter.Reverse("User.UpdateUsername", args).URL
}

func (_ tUser) UpdatePwd(
		oldPwd string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "oldPwd", oldPwd)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("User.UpdatePwd", args).URL
}

func (_ tUser) UpdateTheme(
		theme string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "theme", theme)
	return revel.MainRouter.Reverse("User.UpdateTheme", args).URL
}

func (_ tUser) SendRegisterEmail(
		content string,
		toEmail string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "content", content)
	revel.Unbind(args, "toEmail", toEmail)
	return revel.MainRouter.Reverse("User.SendRegisterEmail", args).URL
}

func (_ tUser) ReSendActiveEmail(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.ReSendActiveEmail", args).URL
}

func (_ tUser) UpdateEmail(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("User.UpdateEmail", args).URL
}

func (_ tUser) ActiveEmail(
		token string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	return revel.MainRouter.Reverse("User.ActiveEmail", args).URL
}

func (_ tUser) UpdateColumnWidth(
		notebookWidth int,
		noteListWidth int,
		mdEditorWidth int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookWidth", notebookWidth)
	revel.Unbind(args, "noteListWidth", noteListWidth)
	revel.Unbind(args, "mdEditorWidth", mdEditorWidth)
	return revel.MainRouter.Reverse("User.UpdateColumnWidth", args).URL
}

func (_ tUser) UpdateLeftIsMin(
		leftIsMin bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "leftIsMin", leftIsMin)
	return revel.MainRouter.Reverse("User.UpdateLeftIsMin", args).URL
}


type tAdminBaseController struct {}
var AdminBaseController tAdminBaseController



type tApiBaseContrller struct {}
var ApiBaseContrller tApiBaseContrller



type tMemberBaseController struct {}
var MemberBaseController tMemberBaseController



type tPreview struct {}
var Preview tPreview


func (_ tPreview) Index(
		userIdOrEmail string,
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("Preview.Index", args).URL
}

func (_ tPreview) Tag(
		userIdOrEmail string,
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("Preview.Tag", args).URL
}

func (_ tPreview) Tags(
		userIdOrEmail string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	return revel.MainRouter.Reverse("Preview.Tags", args).URL
}

func (_ tPreview) Archives(
		userIdOrEmail string,
		notebookId string,
		year int,
		month int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "year", year)
	revel.Unbind(args, "month", month)
	return revel.MainRouter.Reverse("Preview.Archives", args).URL
}

func (_ tPreview) Cate(
		userIdOrEmail string,
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("Preview.Cate", args).URL
}

func (_ tPreview) Post(
		userIdOrEmail string,
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("Preview.Post", args).URL
}

func (_ tPreview) Single(
		userIdOrEmail string,
		singleId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "singleId", singleId)
	return revel.MainRouter.Reverse("Preview.Single", args).URL
}

func (_ tPreview) Search(
		userIdOrEmail string,
		keywords string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userIdOrEmail", userIdOrEmail)
	revel.Unbind(args, "keywords", keywords)
	return revel.MainRouter.Reverse("Preview.Search", args).URL
}


type tAdminBlog struct {}
var AdminBlog tAdminBlog


func (_ tAdminBlog) Index(
		sorter string,
		keywords string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sorter", sorter)
	revel.Unbind(args, "keywords", keywords)
	return revel.MainRouter.Reverse("AdminBlog.Index", args).URL
}

func (_ tAdminBlog) SetRecommend(
		noteId string,
		recommend bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "recommend", recommend)
	return revel.MainRouter.Reverse("AdminBlog.SetRecommend", args).URL
}


type tAdminData struct {}
var AdminData tAdminData


func (_ tAdminData) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminData.Index", args).URL
}

func (_ tAdminData) Backup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminData.Backup", args).URL
}

func (_ tAdminData) Restore(
		createdTime string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "createdTime", createdTime)
	return revel.MainRouter.Reverse("AdminData.Restore", args).URL
}

func (_ tAdminData) Delete(
		createdTime string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "createdTime", createdTime)
	return revel.MainRouter.Reverse("AdminData.Delete", args).URL
}

func (_ tAdminData) UpdateRemark(
		createdTime string,
		remark string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "createdTime", createdTime)
	revel.Unbind(args, "remark", remark)
	return revel.MainRouter.Reverse("AdminData.UpdateRemark", args).URL
}

func (_ tAdminData) Download(
		createdTime string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "createdTime", createdTime)
	return revel.MainRouter.Reverse("AdminData.Download", args).URL
}


type tAdminUpgrade struct {}
var AdminUpgrade tAdminUpgrade


func (_ tAdminUpgrade) UpgradeBlog(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminUpgrade.UpgradeBlog", args).URL
}

func (_ tAdminUpgrade) UpgradeBetaToBeta2(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminUpgrade.UpgradeBetaToBeta2", args).URL
}

func (_ tAdminUpgrade) UpgradeBeta3ToBeta4(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminUpgrade.UpgradeBeta3ToBeta4", args).URL
}


type tAdminUser struct {}
var AdminUser tAdminUser


func (_ tAdminUser) Index(
		sorter string,
		keywords string,
		pageSize int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sorter", sorter)
	revel.Unbind(args, "keywords", keywords)
	revel.Unbind(args, "pageSize", pageSize)
	return revel.MainRouter.Reverse("AdminUser.Index", args).URL
}

func (_ tAdminUser) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminUser.Add", args).URL
}

func (_ tAdminUser) Register(
		email string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("AdminUser.Register", args).URL
}

func (_ tAdminUser) ResetPwd(
		userId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("AdminUser.ResetPwd", args).URL
}

func (_ tAdminUser) DoResetPwd(
		userId string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userId", userId)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("AdminUser.DoResetPwd", args).URL
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Index", args).URL
}

func (_ tAdmin) T(
		t string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "t", t)
	return revel.MainRouter.Reverse("Admin.T", args).URL
}

func (_ tAdmin) GetView(
		view string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "view", view)
	return revel.MainRouter.Reverse("Admin.GetView", args).URL
}


type tAdminEmail struct {}
var AdminEmail tAdminEmail


func (_ tAdminEmail) Email(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminEmail.Email", args).URL
}

func (_ tAdminEmail) Blog(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminEmail.Blog", args).URL
}

func (_ tAdminEmail) DoBlogTag(
		recommendTags string,
		newTags string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "recommendTags", recommendTags)
	revel.Unbind(args, "newTags", newTags)
	return revel.MainRouter.Reverse("AdminEmail.DoBlogTag", args).URL
}

func (_ tAdminEmail) Demo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminEmail.Demo", args).URL
}

func (_ tAdminEmail) DoDemo(
		demoUsername string,
		demoPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "demoUsername", demoUsername)
	revel.Unbind(args, "demoPassword", demoPassword)
	return revel.MainRouter.Reverse("AdminEmail.DoDemo", args).URL
}

func (_ tAdminEmail) ToImage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminEmail.ToImage", args).URL
}

func (_ tAdminEmail) DoToImage(
		toImageBinPath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "toImageBinPath", toImageBinPath)
	return revel.MainRouter.Reverse("AdminEmail.DoToImage", args).URL
}

func (_ tAdminEmail) Set(
		emailHost string,
		emailPort string,
		emailUsername string,
		emailPassword string,
		emailSSL string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "emailHost", emailHost)
	revel.Unbind(args, "emailPort", emailPort)
	revel.Unbind(args, "emailUsername", emailUsername)
	revel.Unbind(args, "emailPassword", emailPassword)
	revel.Unbind(args, "emailSSL", emailSSL)
	return revel.MainRouter.Reverse("AdminEmail.Set", args).URL
}

func (_ tAdminEmail) Template(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminEmail.Template", args).URL
}

func (_ tAdminEmail) SendEmailToEmails(
		sendEmails string,
		latestEmailSubject string,
		latestEmailBody string,
		verified bool,
		saveAsOldEmail bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sendEmails", sendEmails)
	revel.Unbind(args, "latestEmailSubject", latestEmailSubject)
	revel.Unbind(args, "latestEmailBody", latestEmailBody)
	revel.Unbind(args, "verified", verified)
	revel.Unbind(args, "saveAsOldEmail", saveAsOldEmail)
	return revel.MainRouter.Reverse("AdminEmail.SendEmailToEmails", args).URL
}

func (_ tAdminEmail) SendToUsers2(
		emails string,
		latestEmailSubject string,
		latestEmailBody string,
		verified bool,
		saveAsOldEmail bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "emails", emails)
	revel.Unbind(args, "latestEmailSubject", latestEmailSubject)
	revel.Unbind(args, "latestEmailBody", latestEmailBody)
	revel.Unbind(args, "verified", verified)
	revel.Unbind(args, "saveAsOldEmail", saveAsOldEmail)
	return revel.MainRouter.Reverse("AdminEmail.SendToUsers2", args).URL
}

func (_ tAdminEmail) SendEmailDialog(
		emails string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "emails", emails)
	return revel.MainRouter.Reverse("AdminEmail.SendEmailDialog", args).URL
}

func (_ tAdminEmail) SendToUsers(
		userFilterEmail string,
		userFilterWhiteList string,
		userFilterBlackList string,
		latestEmailSubject string,
		latestEmailBody string,
		verified bool,
		saveAsOldEmail bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userFilterEmail", userFilterEmail)
	revel.Unbind(args, "userFilterWhiteList", userFilterWhiteList)
	revel.Unbind(args, "userFilterBlackList", userFilterBlackList)
	revel.Unbind(args, "latestEmailSubject", latestEmailSubject)
	revel.Unbind(args, "latestEmailBody", latestEmailBody)
	revel.Unbind(args, "verified", verified)
	revel.Unbind(args, "saveAsOldEmail", saveAsOldEmail)
	return revel.MainRouter.Reverse("AdminEmail.SendToUsers", args).URL
}

func (_ tAdminEmail) DeleteEmails(
		ids string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ids", ids)
	return revel.MainRouter.Reverse("AdminEmail.DeleteEmails", args).URL
}

func (_ tAdminEmail) List(
		sorter string,
		keywords string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sorter", sorter)
	revel.Unbind(args, "keywords", keywords)
	return revel.MainRouter.Reverse("AdminEmail.List", args).URL
}


type tAdminSetting struct {}
var AdminSetting tAdminSetting


func (_ tAdminSetting) Email(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminSetting.Email", args).URL
}

func (_ tAdminSetting) Blog(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminSetting.Blog", args).URL
}

func (_ tAdminSetting) DoBlogTag(
		recommendTags string,
		newTags string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "recommendTags", recommendTags)
	revel.Unbind(args, "newTags", newTags)
	return revel.MainRouter.Reverse("AdminSetting.DoBlogTag", args).URL
}

func (_ tAdminSetting) ShareNote(
		registerSharedUserId string,
		registerSharedNotebookPerms []int,
		registerSharedNotePerms []int,
		registerSharedNotebookIds []string,
		registerSharedNoteIds []string,
		registerCopyNoteIds []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "registerSharedUserId", registerSharedUserId)
	revel.Unbind(args, "registerSharedNotebookPerms", registerSharedNotebookPerms)
	revel.Unbind(args, "registerSharedNotePerms", registerSharedNotePerms)
	revel.Unbind(args, "registerSharedNotebookIds", registerSharedNotebookIds)
	revel.Unbind(args, "registerSharedNoteIds", registerSharedNoteIds)
	revel.Unbind(args, "registerCopyNoteIds", registerCopyNoteIds)
	return revel.MainRouter.Reverse("AdminSetting.ShareNote", args).URL
}

func (_ tAdminSetting) Demo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminSetting.Demo", args).URL
}

func (_ tAdminSetting) DoDemo(
		demoUsername string,
		demoPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "demoUsername", demoUsername)
	revel.Unbind(args, "demoPassword", demoPassword)
	return revel.MainRouter.Reverse("AdminSetting.DoDemo", args).URL
}

func (_ tAdminSetting) ExportPdf(
		path string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "path", path)
	return revel.MainRouter.Reverse("AdminSetting.ExportPdf", args).URL
}

func (_ tAdminSetting) DoSiteUrl(
		siteUrl string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "siteUrl", siteUrl)
	return revel.MainRouter.Reverse("AdminSetting.DoSiteUrl", args).URL
}

func (_ tAdminSetting) SubDomain(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminSetting.SubDomain", args).URL
}

func (_ tAdminSetting) DoSubDomain(
		noteSubDomain string,
		blogSubDomain string,
		leaSubDomain string,
		blackSubDomains string,
		allowCustomDomain string,
		blackCustomDomains string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteSubDomain", noteSubDomain)
	revel.Unbind(args, "blogSubDomain", blogSubDomain)
	revel.Unbind(args, "leaSubDomain", leaSubDomain)
	revel.Unbind(args, "blackSubDomains", blackSubDomains)
	revel.Unbind(args, "allowCustomDomain", allowCustomDomain)
	revel.Unbind(args, "blackCustomDomains", blackCustomDomains)
	return revel.MainRouter.Reverse("AdminSetting.DoSubDomain", args).URL
}

func (_ tAdminSetting) OpenRegister(
		openRegister string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "openRegister", openRegister)
	return revel.MainRouter.Reverse("AdminSetting.OpenRegister", args).URL
}

func (_ tAdminSetting) HomePage(
		homePage string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "homePage", homePage)
	return revel.MainRouter.Reverse("AdminSetting.HomePage", args).URL
}

func (_ tAdminSetting) Mongodb(
		mongodumpPath string,
		mongorestorePath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "mongodumpPath", mongodumpPath)
	revel.Unbind(args, "mongorestorePath", mongorestorePath)
	return revel.MainRouter.Reverse("AdminSetting.Mongodb", args).URL
}

func (_ tAdminSetting) UploadSize(
		uploadImageSize float64,
		uploadAvatarSize float64,
		uploadBlogLogoSize float64,
		uploadAttachSize float64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uploadImageSize", uploadImageSize)
	revel.Unbind(args, "uploadAvatarSize", uploadAvatarSize)
	revel.Unbind(args, "uploadBlogLogoSize", uploadBlogLogoSize)
	revel.Unbind(args, "uploadAttachSize", uploadAttachSize)
	return revel.MainRouter.Reverse("AdminSetting.UploadSize", args).URL
}


type tApiTag struct {}
var ApiTag tApiTag


func (_ tApiTag) GetSyncTags(
		afterUsn int,
		maxEntry int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "afterUsn", afterUsn)
	revel.Unbind(args, "maxEntry", maxEntry)
	return revel.MainRouter.Reverse("ApiTag.GetSyncTags", args).URL
}

func (_ tApiTag) AddTag(
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("ApiTag.AddTag", args).URL
}

func (_ tApiTag) DeleteTag(
		tag string,
		usn int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	revel.Unbind(args, "usn", usn)
	return revel.MainRouter.Reverse("ApiTag.DeleteTag", args).URL
}


type tApiUser struct {}
var ApiUser tApiUser


func (_ tApiUser) Info(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiUser.Info", args).URL
}

func (_ tApiUser) UpdateUsername(
		username string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	return revel.MainRouter.Reverse("ApiUser.UpdateUsername", args).URL
}

func (_ tApiUser) UpdatePwd(
		oldPwd string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "oldPwd", oldPwd)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("ApiUser.UpdatePwd", args).URL
}

func (_ tApiUser) GetSyncState(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiUser.GetSyncState", args).URL
}

func (_ tApiUser) UpdateLogo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiUser.UpdateLogo", args).URL
}


type tApiAuth struct {}
var ApiAuth tApiAuth


func (_ tApiAuth) Login(
		email string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("ApiAuth.Login", args).URL
}

func (_ tApiAuth) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiAuth.Logout", args).URL
}

func (_ tApiAuth) Register(
		email string,
		pwd string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "pwd", pwd)
	return revel.MainRouter.Reverse("ApiAuth.Register", args).URL
}


type tApiFile struct {}
var ApiFile tApiFile


func (_ tApiFile) GetImage(
		fileId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fileId", fileId)
	return revel.MainRouter.Reverse("ApiFile.GetImage", args).URL
}

func (_ tApiFile) GetAttach(
		fileId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "fileId", fileId)
	return revel.MainRouter.Reverse("ApiFile.GetAttach", args).URL
}

func (_ tApiFile) GetAllAttachs(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("ApiFile.GetAllAttachs", args).URL
}


type tApiNotebook struct {}
var ApiNotebook tApiNotebook


func (_ tApiNotebook) GetSyncNotebooks(
		afterUsn int,
		maxEntry int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "afterUsn", afterUsn)
	revel.Unbind(args, "maxEntry", maxEntry)
	return revel.MainRouter.Reverse("ApiNotebook.GetSyncNotebooks", args).URL
}

func (_ tApiNotebook) GetNotebooks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiNotebook.GetNotebooks", args).URL
}

func (_ tApiNotebook) AddNotebook(
		title string,
		parentNotebookId string,
		seq int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "parentNotebookId", parentNotebookId)
	revel.Unbind(args, "seq", seq)
	return revel.MainRouter.Reverse("ApiNotebook.AddNotebook", args).URL
}

func (_ tApiNotebook) UpdateNotebook(
		notebookId string,
		title string,
		parentNotebookId string,
		seq int,
		usn int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "parentNotebookId", parentNotebookId)
	revel.Unbind(args, "seq", seq)
	revel.Unbind(args, "usn", usn)
	return revel.MainRouter.Reverse("ApiNotebook.UpdateNotebook", args).URL
}

func (_ tApiNotebook) DeleteNotebook(
		notebookId string,
		usn int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	revel.Unbind(args, "usn", usn)
	return revel.MainRouter.Reverse("ApiNotebook.DeleteNotebook", args).URL
}


type tApiNote struct {}
var ApiNote tApiNote


func (_ tApiNote) GetSyncNotes(
		afterUsn int,
		maxEntry int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "afterUsn", afterUsn)
	revel.Unbind(args, "maxEntry", maxEntry)
	return revel.MainRouter.Reverse("ApiNote.GetSyncNotes", args).URL
}

func (_ tApiNote) GetNotes(
		notebookId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "notebookId", notebookId)
	return revel.MainRouter.Reverse("ApiNote.GetNotes", args).URL
}

func (_ tApiNote) GetTrashNotes(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ApiNote.GetTrashNotes", args).URL
}

func (_ tApiNote) GetNote(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("ApiNote.GetNote", args).URL
}

func (_ tApiNote) GetNoteAndContent(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("ApiNote.GetNoteAndContent", args).URL
}

func (_ tApiNote) GetNoteContent(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("ApiNote.GetNoteContent", args).URL
}

func (_ tApiNote) AddNote(
		noteOrContent interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteOrContent", noteOrContent)
	return revel.MainRouter.Reverse("ApiNote.AddNote", args).URL
}

func (_ tApiNote) UpdateNote(
		noteOrContent interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteOrContent", noteOrContent)
	return revel.MainRouter.Reverse("ApiNote.UpdateNote", args).URL
}

func (_ tApiNote) DeleteTrash(
		noteId string,
		usn int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "usn", usn)
	return revel.MainRouter.Reverse("ApiNote.DeleteTrash", args).URL
}

func (_ tApiNote) ExportPdf(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("ApiNote.ExportPdf", args).URL
}


type tMemberIndex struct {}
var MemberIndex tMemberIndex


func (_ tMemberIndex) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberIndex.Index", args).URL
}

func (_ tMemberIndex) T(
		t string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "t", t)
	return revel.MainRouter.Reverse("MemberIndex.T", args).URL
}

func (_ tMemberIndex) GetView(
		view string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "view", view)
	return revel.MainRouter.Reverse("MemberIndex.GetView", args).URL
}


type tMemberUser struct {}
var MemberUser tMemberUser


func (_ tMemberUser) Username(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberUser.Username", args).URL
}

func (_ tMemberUser) Email(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberUser.Email", args).URL
}

func (_ tMemberUser) Password(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberUser.Password", args).URL
}

func (_ tMemberUser) Avatar(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberUser.Avatar", args).URL
}


type tMemberBlog struct {}
var MemberBlog tMemberBlog


func (_ tMemberBlog) Index(
		sorter string,
		keywords string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "sorter", sorter)
	revel.Unbind(args, "keywords", keywords)
	return revel.MainRouter.Reverse("MemberBlog.Index", args).URL
}

func (_ tMemberBlog) UpdateBlogUrlTitle(
		noteId string,
		urlTitle string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "urlTitle", urlTitle)
	return revel.MainRouter.Reverse("MemberBlog.UpdateBlogUrlTitle", args).URL
}

func (_ tMemberBlog) UpdateBlogAbstract(
		noteId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	return revel.MainRouter.Reverse("MemberBlog.UpdateBlogAbstract", args).URL
}

func (_ tMemberBlog) DoUpdateBlogAbstract(
		noteId string,
		imgSrc string,
		desc string,
		abstract string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noteId", noteId)
	revel.Unbind(args, "imgSrc", imgSrc)
	revel.Unbind(args, "desc", desc)
	revel.Unbind(args, "abstract", abstract)
	return revel.MainRouter.Reverse("MemberBlog.DoUpdateBlogAbstract", args).URL
}

func (_ tMemberBlog) Base(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Base", args).URL
}

func (_ tMemberBlog) Comment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Comment", args).URL
}

func (_ tMemberBlog) Paging(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Paging", args).URL
}

func (_ tMemberBlog) Cate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Cate", args).URL
}

func (_ tMemberBlog) UpateCateIds(
		cateIds []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "cateIds", cateIds)
	return revel.MainRouter.Reverse("MemberBlog.UpateCateIds", args).URL
}

func (_ tMemberBlog) UpdateCateUrlTitle(
		cateId string,
		urlTitle string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "cateId", cateId)
	revel.Unbind(args, "urlTitle", urlTitle)
	return revel.MainRouter.Reverse("MemberBlog.UpdateCateUrlTitle", args).URL
}

func (_ tMemberBlog) DoAddOrUpdateSingle(
		singleId string,
		title string,
		content string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "singleId", singleId)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "content", content)
	return revel.MainRouter.Reverse("MemberBlog.DoAddOrUpdateSingle", args).URL
}

func (_ tMemberBlog) AddOrUpdateSingle(
		singleId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "singleId", singleId)
	return revel.MainRouter.Reverse("MemberBlog.AddOrUpdateSingle", args).URL
}

func (_ tMemberBlog) SortSingles(
		singleIds []string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "singleIds", singleIds)
	return revel.MainRouter.Reverse("MemberBlog.SortSingles", args).URL
}

func (_ tMemberBlog) DeleteSingle(
		singleId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "singleId", singleId)
	return revel.MainRouter.Reverse("MemberBlog.DeleteSingle", args).URL
}

func (_ tMemberBlog) UpdateSingleUrlTitle(
		singleId string,
		urlTitle string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "singleId", singleId)
	revel.Unbind(args, "urlTitle", urlTitle)
	return revel.MainRouter.Reverse("MemberBlog.UpdateSingleUrlTitle", args).URL
}

func (_ tMemberBlog) Single(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Single", args).URL
}

func (_ tMemberBlog) Theme(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.Theme", args).URL
}

func (_ tMemberBlog) UpdateTheme(
		themeId string,
		isNew int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	revel.Unbind(args, "isNew", isNew)
	return revel.MainRouter.Reverse("MemberBlog.UpdateTheme", args).URL
}

func (_ tMemberBlog) GetTplContent(
		themeId string,
		filename string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	revel.Unbind(args, "filename", filename)
	return revel.MainRouter.Reverse("MemberBlog.GetTplContent", args).URL
}

func (_ tMemberBlog) UpdateTplContent(
		themeId string,
		filename string,
		content string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	revel.Unbind(args, "filename", filename)
	revel.Unbind(args, "content", content)
	return revel.MainRouter.Reverse("MemberBlog.UpdateTplContent", args).URL
}

func (_ tMemberBlog) DeleteTpl(
		themeId string,
		filename string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	revel.Unbind(args, "filename", filename)
	return revel.MainRouter.Reverse("MemberBlog.DeleteTpl", args).URL
}

func (_ tMemberBlog) ListThemeImages(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.ListThemeImages", args).URL
}

func (_ tMemberBlog) DeleteThemeImage(
		themeId string,
		filename string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	revel.Unbind(args, "filename", filename)
	return revel.MainRouter.Reverse("MemberBlog.DeleteThemeImage", args).URL
}

func (_ tMemberBlog) UploadThemeImage(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.UploadThemeImage", args).URL
}

func (_ tMemberBlog) ActiveTheme(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.ActiveTheme", args).URL
}

func (_ tMemberBlog) DeleteTheme(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.DeleteTheme", args).URL
}

func (_ tMemberBlog) PublicTheme(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.PublicTheme", args).URL
}

func (_ tMemberBlog) ExportTheme(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.ExportTheme", args).URL
}

func (_ tMemberBlog) ImportTheme(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.ImportTheme", args).URL
}

func (_ tMemberBlog) InstallTheme(
		themeId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "themeId", themeId)
	return revel.MainRouter.Reverse("MemberBlog.InstallTheme", args).URL
}

func (_ tMemberBlog) NewTheme(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberBlog.NewTheme", args).URL
}

func (_ tMemberBlog) SetUserBlogBase(
		userBlog interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userBlog", userBlog)
	return revel.MainRouter.Reverse("MemberBlog.SetUserBlogBase", args).URL
}

func (_ tMemberBlog) SetUserBlogComment(
		userBlog interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userBlog", userBlog)
	return revel.MainRouter.Reverse("MemberBlog.SetUserBlogComment", args).URL
}

func (_ tMemberBlog) SetUserBlogStyle(
		userBlog interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "userBlog", userBlog)
	return revel.MainRouter.Reverse("MemberBlog.SetUserBlogStyle", args).URL
}

func (_ tMemberBlog) SetUserBlogPaging(
		perPageSize int,
		sortField string,
		isAsc bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "perPageSize", perPageSize)
	revel.Unbind(args, "sortField", sortField)
	revel.Unbind(args, "isAsc", isAsc)
	return revel.MainRouter.Reverse("MemberBlog.SetUserBlogPaging", args).URL
}


type tMemberGroup struct {}
var MemberGroup tMemberGroup


func (_ tMemberGroup) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MemberGroup.Index", args).URL
}

func (_ tMemberGroup) AddGroup(
		title string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	return revel.MainRouter.Reverse("MemberGroup.AddGroup", args).URL
}

func (_ tMemberGroup) UpdateGroupTitle(
		groupId string,
		title string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "title", title)
	return revel.MainRouter.Reverse("MemberGroup.UpdateGroupTitle", args).URL
}

func (_ tMemberGroup) DeleteGroup(
		groupId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	return revel.MainRouter.Reverse("MemberGroup.DeleteGroup", args).URL
}

func (_ tMemberGroup) AddUser(
		groupId string,
		email string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "email", email)
	return revel.MainRouter.Reverse("MemberGroup.AddUser", args).URL
}

func (_ tMemberGroup) DeleteUser(
		groupId string,
		userId string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "groupId", groupId)
	revel.Unbind(args, "userId", userId)
	return revel.MainRouter.Reverse("MemberGroup.DeleteUser", args).URL
}


