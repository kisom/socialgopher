package main

import (
	"bytes"
	"encoding/json"
	"github.com/gokyle/adn"
)

// NewJSONRequest creates a skeleton request that should
// only need to be filled in with the body.
func (p *Profile) NewJSONRequest() (req *adn.Request) {
	req = new(adn.Request)
	req.Token = p.Identity.Secret()
	req.BodyType = BodyTypeJSON
	return
}

////////////////////
// Post Functions //
////////////////////

// Create a new post. Hashtags are parsed on ADN's end.
func (p *Profile) CreatePost(text string) (post *adn.Post, err error) {
	var args adn.EpArgs
	req := p.NewJSONRequest()

	var body struct {
		Text string `json:"text"`
	}
	body.Text = text

	jBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	req.Body = bytes.NewBuffer(jBody)

	post = new(adn.Post)
	err = p.App.Do(req, "create post", args, post)
	return
}

// Retrieve the user's personalised stream.
func (p *Profile) GetPosts() (posts []*adn.Post, err error) {
	var args adn.EpArgs
	req := p.NewJSONRequest()
	posts = make([]*adn.Post, 0)

	err = p.App.Do(req, "retrieve user personalized stream", args, &posts)
	return
}

// Reply to a post.
func (p *Profile) ReplyTo(postid string, text string) (post *adn.Post, err error) {
	var args adn.EpArgs
	req := p.NewJSONRequest()

	var body struct {
		Text    string `json:"text"`
		ReplyTo string `json:"reply_to"`
	}
	body.Text = text
	body.ReplyTo = postid

	jBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	req.Body = bytes.NewBuffer(jBody)

	post = new(adn.Post)
	err = p.App.Do(req, "create post", args, post)
	return
}

// Attempt to grab a list of all the threads in a post.
func (p *Profile) GetThread(postid string) (thread []*adn.Post, err error) {
	var args adn.EpArgs
	args.Post = postid

	req := p.NewJSONRequest()
	thread = make([]*adn.Post, 0)
	err = p.App.Do(req, "retrieve post replies", args, &thread)
	return
}
