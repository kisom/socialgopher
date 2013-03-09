package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

const (
	image_profile_file = "assets/images/profile_image.png"
	image_stream_file = "assets/images/home.png"
	image_mentions_file = "assets/images/mentions.png"
	image_interactions_file = "assets/images/interactions.png"
	image_stars_file = "assets/images/stars.png"
	image_messages_file = "assets/images/messages.png"
	image_settings_file = "assets/images/settings.png"
)

var (
        assets  map[string]string
)

func init() {
        assets = make(map[string]string, 0)
        assets["profile"] = "assets/images/profile_image.png"
        assets["stream"] = "assets/images/home.png"
        assets["mentions"] = "assets/images/mentions.png"
        assets["interactions"] = "assets/images/interactions.png"
        assets["stars"] = "assets/images/stars.png"
        assets["messages"] = "assets/images/messages.png"
        assets["settings"] = "assets/images/settings.png"
}

func accountWindow() {
	// window settings
	window_account := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window_account.SetPosition(gtk.WIN_POS_CENTER)
	window_account.SetTitle("Add Account")

	// main container 
	container_main := gtk.NewVBox(false, 10)
	container_user := gtk.NewHBox(false, 0) 
	container_pass := gtk.NewHBox(false, 0)
	container_buttons := gtk.NewHBox(false, 5)
	container_main.SetBorderWidth(10)

	// username
	user_label := gtk.NewLabel("Username")
	user_entry := gtk.NewEntry()

	// password
	pass_label := gtk.NewLabel("Password")
	pass_entry := gtk.NewEntry()
	pass_entry.SetVisibility(false)

	// login and cancel buttons
	button_login := gtk.NewButtonWithLabel("Add")
	button_cancel := gtk.NewButtonWithLabel("Cancel")

	// login
	button_login.Clicked(func() {
		// validation holder
		if (user_entry.GetText() == "user" && pass_entry.GetText() == "pass") {
			println("[*] Login successful")
			window_account.Destroy()
		}
	})

	// cancel
	button_cancel.Clicked(func() {
		window_account.Destroy()
	})

	// add elements to containers
	container_buttons.Add(button_login)
	container_buttons.Add(button_cancel)
	container_user.PackStart(user_label, false, false, 20)
	container_user.PackEnd(user_entry, true, true, 1)
	container_pass.PackStart(pass_label, false, false, 20)
	container_pass.PackEnd(pass_entry, true, true, 1)
	container_main.PackStart(container_user, false, false, 1)
	container_main.PackStart(container_pass,  false, false, 1)
	container_main.PackStart(container_buttons,  false, false, 1)

	window_account.Add(container_main)
	window_account.SetSizeRequest(350,150)
	window_account.SetResizable(false)
	window_account.ShowAll()
}

func mainWindow() {
	gtk.Init(&os.Args)

	// window settings
	window_main := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window_main.SetPosition(gtk.WIN_POS_CENTER)
	window_main.SetTitle("Social Gopher")
	window_main.Connect("destroy", func() {
		println("[!] Quit application")
		gtk.MainQuit()
	})

	// images
	image_profile := loadImageAsset("profile")
	image_stream := loadImageAsset("stream")
	image_mentions := loadImageAsset("mentions")
	image_interactions := loadImageAsset("interactions")
	image_stars := loadImageAsset("stars")
	image_messages := loadImageAsset("messages")
	image_settings := loadImageAsset("settings")

	// containers 
	container_main := gtk.NewHBox(false, 1)
	container_left := gtk.NewVBox(false, 1)
	container_right := gtk.NewVBox(false, 5)
	container_compose := gtk.NewHBox(false, 5)
	container_profile := gtk.NewHBox(false, 5)
	container_profile.Add(image_profile)
	container_left.SetBorderWidth(5)
	container_right.SetBorderWidth(5)

	// toolbar
	button_stream := gtk.NewToolButton(image_stream, "My Stream")
	button_mentions := gtk.NewToolButton(image_mentions, "Mentions")
	button_interactions := gtk.NewToolButton(image_interactions, "Interactions")
	button_stars := gtk.NewToolButton(image_stars, "Stars")
	button_messages := gtk.NewToolButton(image_messages, "Messages")
	button_settings := gtk.NewToolButton(image_settings, "Settings")
	button_separator := gtk.NewSeparatorToolItem()
	toolbar := gtk.NewToolbar()
	toolbar.SetOrientation(gtk.ORIENTATION_VERTICAL)
	toolbar.Insert(button_stream, -1)
	toolbar.Insert(button_mentions, -1)
	toolbar.Insert(button_interactions, -1)
	toolbar.Insert(button_stars, -1)
	toolbar.Insert(button_messages, -1)
	toolbar.Insert(button_separator, -1)
	toolbar.Insert(button_settings, -1)

	// stream list
	list_swin := gtk.NewScrolledWindow(nil, nil)
	list_swin.SetPolicy(-1, 1)
	list_swin.SetShadowType(2)
	list_textView := gtk.NewTextView()
	list_textView.SetEditable(false)
	list_textView.SetCursorVisible(false)
	list_textView.SetWrapMode(2)
	list_swin.Add(list_textView)
	list_buffer := list_textView.GetBuffer()

	// compose message
	compose := gtk.NewTextView()
	compose.SetEditable(true)
	compose.SetWrapMode(2)
	compose_swin := gtk.NewScrolledWindow(nil, nil)
	compose_swin.SetPolicy(1, 1)
	compose_swin.SetShadowType(1)
	compose_swin.Add(compose)
	compose_counter := gtk.NewLabel("256")
	compose_buffer := compose.GetBuffer()

	compose_buffer.Connect("changed", func() {
		chars_left := 256 - compose_buffer.GetCharCount()
		compose_counter.SetText(strconv.Itoa(chars_left))
	})

	// post button and counter
	button_post := gtk.NewButtonWithLabel("Post")
	container_post := gtk.NewVBox(false, 1)
	container_post.Add(compose_counter)
	container_post.Add(button_post)

	// button functions
	button_stream.OnClicked(func() {
		list_buffer.SetText("My Stream")
	})
	button_mentions.OnClicked(func() {
		list_buffer.SetText("Mentions")
	})
	button_interactions.OnClicked(func() {
		list_buffer.SetText("Interactions")
	})
	button_stars.OnClicked(func() {
		list_buffer.SetText("Stars")
	})
	button_messages.OnClicked(func() {
		list_buffer.SetText("Messages")
	})
	button_settings.OnClicked(func() {
		accountWindow()
	})
	button_post.Clicked(func() {
		compose_buffer.SetText("")
	})

	// add elements to containers
	container_left.PackStart(container_profile, false, true, 1)
	container_left.PackEnd(toolbar, true, true, 1)
	container_right.PackStart(list_swin, true, true, 1)
	container_right.PackEnd(container_compose, false, false, 1)
	container_compose.PackStart(compose_swin, true, true, 1)
	container_compose.PackEnd(container_post, false, true, 1)
	container_main.PackStart(container_left, false, true, 1)
	container_main.PackEnd(container_right, true, true, 1)

	window_main.Add(container_main)
	window_main.SetSizeRequest(500, 600)
	window_main.ShowAll()

	gtk.Main()
}

func loadImageAsset(assetName string) gtk.IWidget {
        assetFile := assets[assetName]
        return gtk.NewImageFromFile(assetFile)
}

func main() {
	mainWindow()
}
