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

func mainWindow() {
	gtk.Init(&os.Args)

	// window settings
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Social Gopher")
	window.Connect("destroy", func() {
		println("[!] Quit application")
		gtk.MainQuit()
	})

	// images
	image_profile := gtk.NewImageFromFile(image_profile_file)
	image_stream := gtk.NewImageFromFile(image_stream_file)
	image_mentions := gtk.NewImageFromFile(image_mentions_file)
	image_interactions := gtk.NewImageFromFile(image_interactions_file)
	image_stars := gtk.NewImageFromFile(image_stars_file)
	image_messages := gtk.NewImageFromFile(image_messages_file)
	image_settings := gtk.NewImageFromFile(image_settings_file)

	// containers 
	container_main := gtk.NewHBox(false, 1)
	container_left := gtk.NewVBox(false, 1)
	container_right := gtk.NewVBox(false, 10)
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
		list_buffer.SetText("Settings")
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

	window.Add(container_main)
	window.SetSizeRequest(500, 600)
	window.ShowAll()

	gtk.Main()
}

func main() {
	mainWindow()
}
