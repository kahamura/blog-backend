package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ka-aki/blog-backend/graph/generated"
	"github.com/ka-aki/blog-backend/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	var article model.Article
	var user model.User
	article.Title = input.Title
	article.Content = input.Content
	article.CreatedAt = time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local)
	user.Name = "test"
	article.User = &user
	return &article, nil
}

func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	var diary model.Diary
	var user model.User
	diary.Title = input.Title
	diary.Content = input.Content
	diary.CreatedAt = time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local)
	user.Name = "testname"
	diary.User = &user
	return &diary, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	return &model.Article{
		ID:        "1",
		Title:     "This is dummy article title",
		User:      &model.User{Name: "Selena Gomez"},
		Content:   "dummy content",
		CreatedAt: time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local),
	}, nil
}

func (r *queryResolver) Diary(ctx context.Context, id string) (*model.Diary, error) {
	return &model.Diary{
		ID:        "1",
		Title:     "This is dummy diary title",
		User:      &model.User{Name: "Selena Gomez"},
		Content:   "dummy diary content",
		CreatedAt: time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local),
	}, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	var articles []*model.Article
	dummyArticle := model.Article{
		ID:        "2",
		Title:     "my dummy article",
		User:      &model.User{Name: "Britney Spears"},
		Content:   "She is so cute",
		CreatedAt: time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local),
	}
	return append(articles, &dummyArticle), nil
}

func (r *queryResolver) Diaries(ctx context.Context) ([]*model.Diary, error) {
	var diaries []*model.Diary
	dummyDiary := model.Diary{
		ID:        "2",
		Title:     "my dummy diary",
		User:      &model.User{Name: "Britney Spears"},
		Content:   "I spend whole time sleeping today.",
		CreatedAt: time.Date(2022, 1, 01, 0, 0, 0, 0, time.Local),
	}
	return append(diaries, &dummyDiary), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
