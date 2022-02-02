package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ka-aki/blog-backend/graph/generated"
	"github.com/ka-aki/blog-backend/graph/model"
	"github.com/ka-aki/blog-backend/internal/articles"
	"github.com/ka-aki/blog-backend/internal/diaries"
	"github.com/ka-aki/blog-backend/internal/users"
	"github.com/ka-aki/blog-backend/package/jwt"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	article := articles.Article{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	articleID := article.Save()

	return &model.Article{
		ID:        strconv.FormatInt(articleID, 10),
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	diary := diaries.Diary{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	diaryID := diary.Save()

	return &model.Diary{
		ID:        strconv.FormatInt(diaryID, 10),
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt,
		UpdatedAt: diary.UpdatedAt,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username:  input.Username,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	dbArticle := articles.GetArticle(id)
	return &model.Article{
		ID:        dbArticle.ID,
		Title:     dbArticle.Title,
		Content:   dbArticle.Content,
		CreatedAt: dbArticle.CreatedAt,
		UpdatedAt: dbArticle.UpdatedAt,
	}, nil
}

func (r *queryResolver) Diary(ctx context.Context, id string) (*model.Diary, error) {
	dbDiary := diaries.GetDiary(id)
	return &model.Diary{
		ID:        dbDiary.ID,
		Title:     dbDiary.Title,
		Content:   dbDiary.Content,
		CreatedAt: dbDiary.CreatedAt,
		UpdatedAt: dbDiary.UpdatedAt,
	}, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	var resultArticles []*model.Article
	var dbArticles []articles.Article

	dbArticles = articles.GetAll()
	for _, article := range dbArticles {
		resultArticles = append(resultArticles, &model.Article{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		})
	}

	return resultArticles, nil
}

func (r *queryResolver) Diaries(ctx context.Context) ([]*model.Diary, error) {
	var resultDiaries []*model.Diary
	var dbDiaries []diaries.Diary

	dbDiaries = diaries.GetAll()
	for _, diary := range dbDiaries {
		resultDiaries = append(resultDiaries, &model.Diary{
			ID:        diary.ID,
			Title:     diary.Title,
			Content:   diary.Content,
			CreatedAt: diary.CreatedAt,
			UpdatedAt: diary.UpdatedAt,
		})
	}

	return resultDiaries, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
