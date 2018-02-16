package qiitago

import (
	"encoding/json"
	"testing"
	"time"
)

var testCommentJson = []byte(`
{
  "body": "# Example",
  "created_at": "2000-01-01T00:00:00+00:00",
  "id": "3391f50c35f953abfc4f",
  "rendered_body": "<h1>Example</h1>",
  "updated_at": "2000-01-01T00:00:00+00:00",
  "user": {
    "description": "Hello, world.",
    "facebook_id": "yaotti",
    "followees_count": 100,
    "followers_count": 200,
    "github_login_name": "yaotti",
    "id": "yaotti",
    "items_count": 300,
    "linkedin_id": "yaotti",
    "location": "Tokyo, Japan",
    "name": "Hiroshige Umino",
    "organization": "Increments Inc",
    "permanent_id": 1,
    "profile_image_url": "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
    "twitter_screen_name": "yaotti",
    "website_url": "http://yaotti.hatenablog.com"
  }
}
`)

var testComment = Comment{
	Id:           "3391f50c35f953abfc4f",
	RenderedBody: "<h1>Example</h1>",
	Body:         "# Example",
	CreatedAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	User: User{
		Id:                "yaotti",
		Description:       &description,
		FacebookId:        &name,
		FolloweesCount:    100,
		FollowersCount:    200,
		GithubLoginName:   &name,
		ItemsCount:        300,
		LinkedinId:        &name,
		Location:          &location,
		Name:              &longName,
		Organization:      &organization,
		PermanentId:       1,
		ProfileImageUrl:   "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
		TwitterScreenName: &name,
		WebsiteUrl:        &websiteUrl,
	},
}

var testTeam = Team{
	Id:     "increments",
	Active: true,
	Name:   "Increments Inc.",
}

var testTeamJson = []byte(`
{
  "active": true,
  "id": "increments",
  "name": "Increments Inc."
}
`)

var testPostsJson = []byte(`
[
  {
    "rendered_body": "<h1>Example</h1>",
    "body": "# Example",
    "coediting": false,
    "comments_count": 100,
    "created_at": "2000-01-01T00:00:00+00:00",
    "group": {
      "created_at": "2000-01-01T00:00:00+00:00",
      "id": 1,
      "name": "Dev",
      "private": false,
      "updated_at": "2000-01-01T00:00:00+00:00",
      "url_name": "dev"
    },
    "id": "4bd431809afb1bb99e4f",
    "likes_count": 100,
    "private": false,
    "reactions_count": 100,
    "tags": [
      {
        "name": "Ruby",
        "versions": [
          "0.0.1"
        ]
      }
    ],
    "title": "Example title",
    "updated_at": "2000-01-01T00:00:00+00:00",
    "url": "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
    "user": {
      "description": "Hello, world.",
      "facebook_id": "yaotti",
      "followees_count": 100,
      "followers_count": 200,
      "github_login_name": "yaotti",
      "id": "yaotti",
      "items_count": 300,
      "linkedin_id": "yaotti",
      "location": "Tokyo, Japan",
      "name": "Hiroshige Umino",
      "organization": "Increments Inc",
      "permanent_id": 1,
      "profile_image_url": "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
      "twitter_screen_name": "yaotti",
      "website_url": "http://yaotti.hatenablog.com"
    },
    "page_views_count": 100
  }
]
`)

var description = "Hello, world."
var name = "yaotti"
var location = "Tokyo, Japan"
var longName = "Hiroshige Umino"
var organization = "Increments Inc"
var websiteUrl = "http://yaotti.hatenablog.com"
var pageViewCount = 100

var testPosts = Posts{
	Post{
		Id:            "4bd431809afb1bb99e4f",
		RenderedBody:  "<h1>Example</h1>",
		Body:          "# Example",
		Coediting:     false,
		CommentsCount: 100,
		CreatedAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Group: &Group{
			Id:        1,
			CreatedAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			Name:      "Dev",
			Private:   false,
			UpdatedAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UrlName:   "dev",
		},
		LikesCount:     100,
		Private:        false,
		ReactionsCount: 100,
		Tags: Taggings{
			Tagging{
				Name: "Ruby",
				Versions: []string{
					"0.0.1",
				},
			},
		},
		Title:     "Example title",
		UpdatedAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Url:       "https://qiita.com/yaotti/items/4bd431809afb1bb99e4f",
		User: User{
			Id:              "yaotti",
			Description:     &description,
			FacebookId:      &name,
			FolloweesCount:  100,
			FollowersCount:  200,
			GithubLoginName: &name,

			ItemsCount:        300,
			LinkedinId:        &name,
			Location:          &location,
			Name:              &longName,
			Organization:      &organization,
			PermanentId:       1,
			ProfileImageUrl:   "https://si0.twimg.com/profile_images/2309761038/1ijg13pfs0dg84sk2y0h_normal.jpeg",
			TwitterScreenName: &name,
			WebsiteUrl:        &websiteUrl,
		},
		PageViewsCount: &pageViewCount,
	},
}

func TestUnmarshalPosts(t *testing.T) {
	ps := Posts{}
	err := json.Unmarshal(testPostsJson, &ps)
	if err != nil {
		t.Fatal(err)
	}
	if !PostValueEqual(&testPosts[0], &ps[0]) {
		t.Fatalf("Unmarshaled not matched.\nwant: %v\nhave: %v\n", testPosts, ps)
	}
}

func TestUnmarshalComment(t *testing.T) {
	c := Comment{}
	err := json.Unmarshal(testCommentJson, &c)
	if err != nil {
		t.Fatal(err)
	}
	if !CommentValueEqual(&testComment, &c) {
		t.Fatalf("Unmarshaled not matched.\nwant: %v\nhave: %v\n", testComment, c)
	}
}

func TestUnmarshalTeam(t *testing.T) {
	team := Team{}
	err := json.Unmarshal(testTeamJson, &team)
	if err != nil {
		t.Fatal(err)
	}
	if testTeam != team {
		t.Fatalf("Unmarshaled not matched.\nwant: %v\nhave: %v\n", testTeam, team)
	}
}
