// Copyright 2018 Yutaka Nishimura. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
   The qiitago package implements a data struct mapper
   for Qiita API v2.

   Official documents: https://qiita.com/api/v2/docs

   Json and go data struct binding is provided.
   Value equal helper function is provided, because of
   support null data and pointors in struct.

   The name of types are same as Qiita API path.
   The types that have Prefix "Post" are that have
   same api path with "GET" and different return data
   structure. So no api exists "post_xxx", but "POST xxx"
   is available in Qiita API.
*/
package qiitago

import (
	"reflect"
	"time"
)

// Comment is struct for "comment" in Qiita api.
type Comment struct {
	Id           string    `json:"id"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	RenderedBody string    `json:"rendered_body"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user"`
}

// Post is struct for "post" in Qiita api.
type Post struct {
	Id             string    `json:"id"`
	RenderedBody   string    `json:"rendered_body"`
	Body           string    `json:"body"`
	Coediting      bool      `json:"coediting"`
	CommentsCount  int       `json:"comments_count"`
	CreatedAt      time.Time `json:"created_at"`
	Group          *Group    `json:"group"`
	LikesCount     int       `json:"likes_count"`
	Private        bool      `json:"private"`
	ReactionsCount int       `json:"reactions_count"`
	Tags           Taggings  `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
	User           User      `json:"user"`
	PageViewsCount *int      `json:"page_views_count"`
}

// Posts is struct for array of post in Qiita api.
type Posts []Post

// User is struct for "user" in Qiita api.
type User struct {
	Id                string  `json:"id"`
	Description       *string `json:"description"`
	FacebookId        *string `json:"facebook_id"`
	FolloweesCount    int     `json:"followees_count"`
	FollowersCount    int     `json:"followers_count"`
	GithubLoginName   *string `json:"github_login_name"`
	ItemsCount        int     `json:"items_count"`
	LinkedinId        *string `json:"linkedin_id"`
	Location          *string `json:"location"`
	Name              *string `json:"name"`
	Organization      *string `json:"organization"`
	PermanentId       int     `json:"permanent_id"`
	ProfileImageUrl   string  `json:"profile_image_url"`
	TwitterScreenName *string `json:"twitter_screen_name"`
	WebsiteUrl        *string `json:"website_url"`
}

// Users is struct for array of "user" in Qiita api.
type Users []User

// Tagging is struct for "tagging" in Qiita api.
type Tagging struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// Taggings is struct for array of "tagging" in Qiita api.
type Taggings []Tagging

// Tag is struct for "tag" in Qiita api.
type Tag struct {
	Id             string  `json:"id"`
	FollowersCount int     `json:"followers_count"`
	IconUrl        *string `json:"icon_url"` //for null
	ItemsCount     int     `json:"items_count"`
}

// Tags is struct for array of "tag" in Qiita api.
type Tags []Tag

// Group is struct for "group" in Qiita api.
type Group struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	UpdatedAt time.Time `json:"updated_at"`
	UrlName   string    `json:"url_name"`
}

// Team is struct for "team" in Qiita api.
type Team struct {
	Id     string `json:"id"`
	Active bool   `json:"active"`
	Name   string `json:"name"`
}

// Teams is struct for array of "team" in Qiita api.
type Teams []Team

// Template is struct for "template" in Qiita api.
type Template struct {
	Id            int      `json:"id"`
	Body          string   `json:"body"`
	Name          string   `json:"name"`
	ExpandedBody  string   `json:"expanded_body"`
	ExpandedTags  Taggings `json:"expanded_tags"`
	ExpandedTitle string   `json:"expanded_title"`
	Tags          Taggings `json:"tags"`
	Title         string   `json:"title"`
}

// Templates is struct for array of "template" in Qiita api.
type Templates []Template

// PostTemplate is struct for POST "template"(create new template) in Qiita api.
type PostTemplate struct {
	Body  string   `json:"body"`
	Name  string   `json:"name"`
	Tags  Taggings `json:"tags"`
	Title string   `json:"title"`
}

// Project is struct for "project" in Qiita api.
type Project struct {
	Id             int       `json:"id"`
	RenderedBody   string    `json:"rendered_body"`
	Archived       bool      `json:"archived"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	ReactionsCount int       `json:"reactions_count"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Projects is type for array of "project" in Qiita api.
type Projects []Project

// PostProject is struct for POST "project"(create new project) in Qiita api.
type PostProject struct {
	Archived bool     `json:"archived"`
	Body     string   `json:"body"`
	Name     string   `json:"name"`
	Tags     Taggings `json:"tags"`
}

// ExpandedTemplate is struct for "expanded_templates" in Qiita api.
type ExpandedTemplate struct {
	Body  string   `json:"body"`
	Tags  Taggings `json:"tags"`
	Title string   `json:"title"`
}

// Reaction is struct for "reaction" in Qiita api.
type Reaction struct {
	CreatedAt time.Time    `json:"created_at"`
	ImageUrl  string       `json:"image_url"`
	Name      ReactionName `json:"name"`
	User      User         `json:"user"`
}

// ReactionName is type for "reaction_name" in Qiita api.
type ReactionName string

// Reactions is type for array of "reaction" in Qiita api.
type Reactions []Reaction

// AuthenticatedUser is struct for "authenticated_user" in Qiita api.
type AuthenticatedUser struct {
	Id                          string  `json:"id"`
	Description                 *string `json:"description"`
	FacebookId                  *string `json:"facebook_id"`
	FolloweesCount              int     `json:"followees_count"`
	FollowersCount              int     `json:"followers_count"`
	GithubLoginName             *string `json:"github_login_name"`
	ItemsCount                  int     `json:"items_count"`
	LinkedinId                  *string `json:"linkedin_id"`
	Location                    *string `json:"location"`
	Name                        *string `json:"name"`
	Organization                *string `json:"organization"`
	PermanentId                 int     `json:"permanent_id"`
	ProfileImageUrl             string  `json:"profile_image_url"`
	TwitterScreenName           *string `json:"twitter_screen_name"`
	WebsiteUrl                  *string `json:"website_url"`
	ImageMonthlyUploadLimit     int     `json:"image_monthly_upload_limit"`
	ImageMonthlyUploadRemaining int     `json:"image_monthly_upload_remaining"`
	TeamOnly                    bool    `json:"team_only"`
}

// UserValueEqual check instance value equality between 2 Users.
func UserValueEqual(u1 *User, u2 *User) bool {
	return u1.Id == u2.Id &&
		*u1.Description == *u2.Description &&
		*u1.FacebookId == *u2.FacebookId &&
		u1.FolloweesCount == u2.FolloweesCount &&
		u1.FollowersCount == u2.FollowersCount &&
		*u1.GithubLoginName == *u2.GithubLoginName &&
		u1.ItemsCount == u2.ItemsCount &&
		*u1.LinkedinId == *u2.LinkedinId &&
		*u1.Location == *u2.Location &&
		*u1.Name == *u2.Name &&
		*u1.Organization == *u2.Organization &&
		u1.PermanentId == u2.PermanentId &&
		u1.ProfileImageUrl == u2.ProfileImageUrl &&
		*u1.TwitterScreenName == *u2.TwitterScreenName &&
		*u1.WebsiteUrl == *u2.WebsiteUrl
}

// GroupValueEqual check instance value equality between 2 Groups.
func GroupValueEqual(g1 *Group, g2 *Group) bool {
	return g1.Id == g2.Id &&
		g1.CreatedAt.Equal(g2.CreatedAt) &&
		g1.Name == g2.Name &&
		g1.Private == g2.Private &&
		g1.UpdatedAt.Equal(g2.UpdatedAt) &&
		g1.UrlName == g2.UrlName
}

// PostValueEqual check instance value equality between 2 Posts.
func PostValueEqual(p1 *Post, p2 *Post) bool {
	return p1.Id == p2.Id &&
		p1.RenderedBody == p2.RenderedBody &&
		p1.Body == p2.Body &&
		p1.Coediting == p2.Coediting &&
		p1.CommentsCount == p2.CommentsCount &&
		p1.CreatedAt.Equal(p2.CreatedAt) &&
		GroupValueEqual(p1.Group, p2.Group) &&
		p1.LikesCount == p2.LikesCount &&
		p1.Private == p2.Private &&
		p1.ReactionsCount == p2.ReactionsCount &&
		reflect.DeepEqual(p1.Tags, p2.Tags) &&
		p1.Title == p2.Title &&
		p1.UpdatedAt.Equal(p2.UpdatedAt) &&
		p1.Url == p2.Url &&
		UserValueEqual(&p1.User, &p2.User) &&
		*p1.PageViewsCount == *p2.PageViewsCount
}

// CommentValueEqual check instance value equality between 2 Comments.
func CommentValueEqual(c1 *Comment, c2 *Comment) bool {
	return c1.Id == c2.Id &&
		c1.Body == c2.Body &&
		c1.CreatedAt.Equal(c2.CreatedAt) &&
		c1.RenderedBody == c2.RenderedBody &&
		c1.UpdatedAt.Equal(c2.UpdatedAt) &&
		UserValueEqual(&c1.User, &c2.User)
}

// ProjectValueEqual check instance value equality between 2 Projects.
func ProjectValueEqual(p1 *Project, p2 *Project) bool {
	return p1.Id == p2.Id &&
		p1.RenderedBody == p2.RenderedBody &&
		p1.Archived == p2.Archived &&
		p1.Body == p2.Body &&
		p1.CreatedAt.Equal(p2.CreatedAt) &&
		p1.Name == p2.Name &&
		p1.ReactionsCount == p2.ReactionsCount &&
		p1.UpdatedAt.Equal(p2.UpdatedAt)
}

// ReactionValueEqual check instance value equality between 2 Reaction.
func ReactionValueEqual(r1 *Reaction, r2 *Reaction) bool {
	return r1.CreatedAt.Equal(r2.CreatedAt) &&
		r1.ImageUrl == r2.ImageUrl &&
		r1.Name == r2.Name &&
		UserValueEqual(&r1.User, &r2.User)
}

// AuthenticatedUserValueEqual check instance value equality between 2 AuthenticatedUser.
func AuthenticatedUserValueEqual(u1 *AuthenticatedUser, u2 *AuthenticatedUser) bool {
	return u1.Id == u2.Id &&
		*u1.Description == *u2.Description &&
		*u1.FacebookId == *u2.FacebookId &&
		u1.FolloweesCount == u2.FolloweesCount &&
		u1.FollowersCount == u2.FollowersCount &&
		*u1.GithubLoginName == *u2.GithubLoginName &&
		u1.ItemsCount == u2.ItemsCount &&
		*u1.LinkedinId == *u2.LinkedinId &&
		*u1.Location == *u2.Location &&
		*u1.Name == *u2.Name &&
		*u1.Organization == *u2.Organization &&
		u1.PermanentId == u2.PermanentId &&
		u1.ProfileImageUrl == u2.ProfileImageUrl &&
		*u1.TwitterScreenName == *u2.TwitterScreenName &&
		*u1.WebsiteUrl == *u2.WebsiteUrl &&
		u1.ImageMonthlyUploadLimit == u2.ImageMonthlyUploadLimit &&
		u1.ImageMonthlyUploadRemaining == u2.ImageMonthlyUploadRemaining &&
		u1.TeamOnly == u1.TeamOnly
}
