package qiitago

import (
	"reflect"
	"time"
)

type Post struct {
	Id             string    `json:"id"`
	RenderedBody   string    `json:"rendered_body"`
	Body           string    `json:"body"`
	Coediting      bool      `json:"coediting"`
	CommentsCount  int       `json:"comments_count"`
	CreatedAt      time.Time `json:"craeated_at"`
	Group          *Group    `json:"group"`
	LikesCount     int       `json:"likes_count"`
	Private        bool      `json:"private"`
	ReactionsCount int       `json:"reactions_count"`
	Tags           PostTags  `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	Url            string    `json:"url"`
	User           User      `json:"user"`
	PageViewsCount *int      `json:"page_views_count"`
}

type User struct {
	Id                string  `json:"id"`
	Description       *string `json:"description"`
	FacebookId        *string `json:"facebook_id"`
	FolloweesCount    int     `json:followees_count`
	FollowersCount    int     `json:followers_count`
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

type PostTag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type PostTags []PostTag

type Tag struct {
	Id             string  `json:"id"`
	FollowersCount int     `json:"followers_count"`
	IconUrl        *string `json:"icon_url"` //for null
	ItemsCount     int     `json:"items_count"`
}

type Group struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	UpdatedAt time.Time `json:"updated_at"`
	UrlName   string    `json:"url_name"`
}

type Tags []Tag

type Posts []Post

func UserValueEqual(u1 *User, u2 *User) bool {
	return u1.Id == u2.Id &&
		&u1.Description == &u2.Description &&
		&u1.FacebookId == &u2.FacebookId &&
		u1.FolloweesCount == u2.FolloweesCount &&
		u1.FollowersCount == u2.FollowersCount &&
		&u1.GithubLoginName == &u1.GithubLoginName &&
		u1.ItemsCount == u2.ItemsCount &&
		&u1.LinkedinId == &u1.LinkedinId &&
		&u1.Location == &u1.Location &&
		&u1.Name == &u1.Name &&
		&u1.Organization == &u1.Organization &&
		u1.PermanentId == u2.PermanentId &&
		u1.ProfileImageUrl == u1.ProfileImageUrl &&
		&u1.TwitterScreenName == &u2.TwitterScreenName &&
		&u1.WebsiteUrl == &u2.WebsiteUrl
}

func PostValueEqual(p1 Post, p2 Post) bool {
	return p1.Id == p2.Id &&
		p1.RenderedBody == p2.RenderedBody &&
		p1.Body == p2.Body &&
		p1.Coediting == p2.Coediting &&
		p1.CommentsCount == p2.CommentsCount &&
		p1.CreatedAt == p2.CreatedAt &&
		reflect.DeepEqual(p1.Group, &p2.Group) &&
		p1.LikesCount == p2.LikesCount &&
		p1.Private == p2.Private &&
		p1.ReactionsCount == p2.ReactionsCount &&
		reflect.DeepEqual(p1.Tags, p2.Tags) &&
		p1.Title == p2.Title &&
		p1.UpdatedAt == p2.UpdatedAt &&
		p1.Url == p2.Url &&
		UserValueEqual(&p1.User, &p2.User) &&
		&p1.PageViewsCount == &p2.PageViewsCount
}
