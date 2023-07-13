// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	"bufio"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/cache"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/utils"
)

const (
	managefileindex__21 = `<div class="files-list" style="">`
	managefileindex__29 = `<script>listenDeleteNodeEvents('file', '/files', '/manage/files')</script></body></html>`
	managefileindex__80 = `"/></a><div class="actions" style="font-size:.86rem"><div><a href="`
	managefileindex__82 = `</a></div><a class="delete-file" href="#" data-id="`
)

func ManageFileIndex(paginate *entities.Paginate[entities.File]) func(meta *entities.Meta, wr *bufio.Writer) {
	return func(meta *entities.Meta, wr *bufio.Writer) {
		buffer := &WriterAsBuffer{wr}

		buffer.WriteString(commentlist__0)

		var title = meta.GetTitle()
		var appName = config.Setting("app_name")
		var appLogo = config.Setting("app_logo")
		buffer.WriteString(commentlist__1)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__2)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__3)
		WriteAll(meta.Type, true, buffer)
		buffer.WriteString(commentlist__4)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__5)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__6)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__7)
		WriteAll(config.Setting("twitter_site"), true, buffer)
		buffer.WriteString(commentlist__8)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__9)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__10)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__11)
		WriteAll(appName+" Feed", true, buffer)
		buffer.WriteString(commentlist__12)
		WriteAll(utils.Url("/feed"), true, buffer)
		buffer.WriteString(commentlist__13)
		if appLogo != "" {
			buffer.WriteString(commentlist__30)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__31)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Description != "" {
			buffer.WriteString(commentlist__33)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__34)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__35)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Image != "" {
			buffer.WriteString(commentlist__37)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__38)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		WriteAll(asset.CssFile("css/light.min.css"), false, buffer)
		WriteAll(asset.CssFile("css/style.css"), false, buffer)
		WriteAll(config.Setting("inject_header"), false, buffer)
		buffer.WriteString(commentlist__14)
		WriteAll(utils.Url(""), true, buffer)
		buffer.WriteString(commentlist__15)
		var logoUrl = config.Setting("app_logo")
		if logoUrl != "" {
			buffer.WriteString(commentlist__40)
			WriteAll(logoUrl, true, buffer)
			buffer.WriteString(commentlist__41)
			WriteAll(config.Setting("app_name"), true, buffer)
			buffer.WriteString(commentlist__13)
		} else {
			buffer.WriteString(commentlist__43)

		}
		buffer.WriteString(commentlist__16)
		WriteAll(meta.Query, true, buffer)
		buffer.WriteString(commentlist__17)
		WriteAll(utils.Url("/search"), true, buffer)
		buffer.WriteString(commentlist__18)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__45)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__46)

		} else {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/posts/new"), true, buffer)
			buffer.WriteString(commentlist__48)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(meta.User.Username, true, buffer)
			buffer.WriteString(commentlist__50)
			if meta.User.AvatarImageUrl != "" {
				buffer.WriteString(commentlist__57)
				WriteAll(meta.User.AvatarImageUrl, true, buffer)
				buffer.WriteString(commentlist__41)
				WriteAll(meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__13)
			} else {
				buffer.WriteString(commentlist__60)

			}
			buffer.WriteString(commentlist__51)

			if meta.User != nil && meta.User.IsRoot() {
				buffer.WriteString(commentlist__44)
				WriteAll(utils.Url("/manage"), true, buffer)
				buffer.WriteString(commentlist__62)

			}
			buffer.WriteString(commentlist__44)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__53)
			WriteAll(utils.Url("/posts"), true, buffer)
			buffer.WriteString(commentlist__54)
			WriteAll(utils.Url("/settings"), true, buffer)
			buffer.WriteString(commentlist__55)
			WriteAll(utils.Url("/logout"), true, buffer)
			buffer.WriteString(commentlist__56)

		}
		buffer.WriteString(commentlist__19)

		{
			buffer.WriteString(commentlist__119)
			WriteAll(utils.Url("/manage"), true, buffer)
			buffer.WriteString(commentlist__120)
			WriteAll(utils.Url("/manage/topics"), true, buffer)
			buffer.WriteString(commentlist__121)
			WriteAll(utils.Url("/manage/posts"), true, buffer)
			buffer.WriteString(commentlist__122)
			WriteAll(utils.Url("/manage/pages"), true, buffer)
			buffer.WriteString(commentlist__123)
			WriteAll(utils.Url("/manage/roles"), true, buffer)
			buffer.WriteString(commentlist__124)
			WriteAll(utils.Url("/manage/users"), true, buffer)
			buffer.WriteString(commentlist__125)
			WriteAll(utils.Url("/manage/comments"), true, buffer)
			buffer.WriteString(commentlist__126)
			WriteAll(utils.Url("/manage/files"), true, buffer)
			buffer.WriteString(commentlist__127)
			WriteAll(utils.Url("/manage/settings"), true, buffer)
			buffer.WriteString(commentlist__72)

		}

		buffer.WriteString(filelist__20)

		{
			var (
				msgs = meta.Messages
			)

			if msgs.Length() > 0 {
				buffer.WriteString(commentlist__73)
				var messages = msgs.Get()
				for _, msg := range messages {
					buffer.WriteString(commentlist__75)
					WriteAll(msg.Type, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(msg.Message, true, buffer)
					buffer.WriteString(commentlist__77)
				}
				buffer.WriteString(commentlist__74)
			}
		}

		buffer.WriteString(managefileindex__21)
		for _, file := range paginate.Data {
			var fileUrl = file.Url()
			buffer.WriteString(filelist__78)
			WriteAll(fileUrl, true, buffer)
			buffer.WriteString(filelist__79)
			WriteAll(fileUrl, true, buffer)
			buffer.WriteString(managefileindex__80)
			WriteAll(file.User.Url(), true, buffer)
			buffer.WriteString(commentlist__87)
			WriteAll(file.User.Name(), true, buffer)
			buffer.WriteString(managefileindex__82)
			WriteAll(file.ID, true, buffer)
			buffer.WriteString(filelist__81)

		}
		buffer.WriteString(commentlist__22)
		var links = paginate.Links()
		buffer.WriteString(commentlist__23)
		for _, link := range links {
			buffer.WriteString(commentlist__44)
			WriteAll(link.Link, true, buffer)
			buffer.WriteString(commentlist__103)
			WriteAll(link.Class, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll(link.Label, true, buffer)
			buffer.WriteString(commentlist__105)

		}
		buffer.WriteString(commentlist__24)
		WriteAll(config.Setting("app_name"), true, buffer)
		buffer.WriteString(commentlist__25)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__106)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__107)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__108)

		} else {
			{
				buffer.WriteString(commentlist__63)
				WriteAll(meta.User.AvatarElm("32", "32", false), false, buffer)
				buffer.WriteString(commentlist__64)
				WriteAll(meta.User.Url(), true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(meta.User.Name(), true, buffer)
				buffer.WriteString(commentlist__66)
				WriteAll("@"+meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__67)
				WriteAll(utils.Url("/posts/new"), true, buffer)
				buffer.WriteString(commentlist__68)
				WriteAll(utils.Url("/posts"), true, buffer)
				buffer.WriteString(commentlist__69)
				WriteAll(utils.Url("/comments"), true, buffer)
				buffer.WriteString(commentlist__70)
				WriteAll(utils.Url("/files"), true, buffer)
				buffer.WriteString(commentlist__71)
				WriteAll(utils.Url("/settings"), true, buffer)
				buffer.WriteString(commentlist__72)

			}

			if meta.User.IsRoot() {
				{
					buffer.WriteString(commentlist__119)
					WriteAll(utils.Url("/manage"), true, buffer)
					buffer.WriteString(commentlist__120)
					WriteAll(utils.Url("/manage/topics"), true, buffer)
					buffer.WriteString(commentlist__121)
					WriteAll(utils.Url("/manage/posts"), true, buffer)
					buffer.WriteString(commentlist__122)
					WriteAll(utils.Url("/manage/pages"), true, buffer)
					buffer.WriteString(commentlist__123)
					WriteAll(utils.Url("/manage/roles"), true, buffer)
					buffer.WriteString(commentlist__124)
					WriteAll(utils.Url("/manage/users"), true, buffer)
					buffer.WriteString(commentlist__125)
					WriteAll(utils.Url("/manage/comments"), true, buffer)
					buffer.WriteString(commentlist__126)
					WriteAll(utils.Url("/manage/files"), true, buffer)
					buffer.WriteString(commentlist__127)
					WriteAll(utils.Url("/manage/settings"), true, buffer)
					buffer.WriteString(commentlist__72)

				}

			}
		}
		buffer.WriteString(commentlist__26)

		for _, topic := range cache.Topics {
			buffer.WriteString(commentlist__106)
			WriteAll(topic.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(topic.Name, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll("#"+topic.Name, true, buffer)
			buffer.WriteString(commentlist__132)
		}
		buffer.WriteString(commentlist__27)
		WriteAll(config.Setting("footer_content"), false, buffer)
		buffer.WriteString(commentlist__28)
		WriteAll(config.Setting("inject_footer"), false, buffer)
		WriteAll(asset.JsFile("js/layout.js"), false, buffer)
		WriteAll(asset.JsFile("js/main.js"), false, buffer)
		buffer.WriteString(managefileindex__29)

	}
}