package postgres

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (db *Postgres) GetUser(ctx context.Context, token string) (*model.UserData, error) {
	/*	span := jaeger.GetSpan(ctx, "GetUser")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select u.id, u.login, u.email, u."token", ur."role"
	from "user" u
		join user_role ur on ur.id = u.role_id
	where u."token" = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "token": token}, "GetUser")

	var res model.UserData
	err := db.GetClient().GetContext(ctx, &res, query, token)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "token": token}, "GetUser error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) UserAuth(ctx context.Context, request model.UserAuthRequest) (*model.UserData, error) {
	/*	span := jaeger.GetSpan(ctx, "UserAuth")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select u.id, u.login, u.email, u."token", ur."role"
	from "user" u
		join user_role ur on ur.id = u.role_id
	where u.login = $1 and u.password = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserAuth")

	var res model.UserData
	err := db.GetClient().GetContext(ctx, &res, query, request.Login, getMD5Hash(request.Password))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserAuth error", err)
		}
		return nil, err
	}

	if res.Token == nil {
		query = `update "user" set token = get_uuid() where id = $1 returning token`

		//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserAuth")

		token := ""
		err = db.GetClient().GetContext(ctx, &token, query, res.Id)
		if err != nil {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserAuth error", err)
			return nil, err
		}

		res.Token = &token
	}

	return &res, nil
}

func (db *Postgres) UserRegister(ctx context.Context, request model.UserRegisterRequest) (*model.UserData, error) {
	/*	span := jaeger.GetSpan(ctx, "UserRegister")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select count(u.id) from "user" u where u.login = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserRegister")

	var cnt int
	err := db.GetClient().GetContext(ctx, &cnt, query, request.Login)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserRegister error", err)
		}
		return nil, err
	}
	if cnt > 0 {
		return nil, nil
	}

	query = `insert into "user" (login, password, email, role_id, token)
	select $1, $2, $3, id, get_uuid() from user_role where role = 'Common'
	returning id, login, email, "token", 'Common' as role`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserRegister")

	var res model.UserData
	err = db.GetClient().GetContext(ctx, &res, query, request.Login, getMD5Hash(request.Password), request.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserRegister error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) UserResetPassword(ctx context.Context, request model.UserResetPasswordRequest, password string) (*model.UserData, error) {
	/*	span := jaeger.GetSpan(ctx, "UserResetPassword")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select u.id, u.login, u.email, u."token", ur."role"
	from "user" u
		join user_role ur on ur.id = u.role_id
	where u.login = $1 and u.email = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserResetPassword")

	var res model.UserData
	err := db.GetClient().GetContext(ctx, &res, query, request.Login, request.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserResetPassword error", err)
		}
		return nil, err
	}

	query = `update "user" set password = $1, token = null where id = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserResetPassword")

	_, err = db.GetClient().ExecContext(ctx, query, getMD5Hash(password), res.Id)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserResetPassword error", err)
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) UserChangeData(ctx context.Context, userId int, request model.UserChangeDataRequest) (*model.UserData, error) {
	/*	span := jaeger.GetSpan(ctx, "UserChangeData")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	args := []interface{}{
		userId,
	}
	i := 2

	query := `update "user" set token = get_uuid() `

	if request.Email != nil && len(*request.Email) > 0 {
		query += fmt.Sprintf(", email = $%v ", i)
		args = append(args, *request.Email)
		i++
	}
	if request.Password != nil && len(*request.Password) > 0 {
		query += fmt.Sprintf(", password = $%v ", i)
		args = append(args, getMD5Hash(*request.Password))
		i++
	}

	query += ` where id = $1
	returning id, login, email, "token", 'Common' as role`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserChangeData")

	var res model.UserData
	err := db.GetClient().GetContext(ctx, &res, query, args...)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "UserChangeData error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) GetUserFavourites(ctx context.Context, userId int) ([]model.Favourites, error) {
	/*	span := jaeger.GetSpan(ctx, "GetUserFavourites")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select uf."name", uf.json_array::jsonb as favourites, uf.guid from user_favourites uf where uf.user_id = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId}, "GetUserFavourites")

	res := make([]model.Favourites, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, userId)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId}, "GetUserFavourites error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) AddUserFavourites(ctx context.Context, userId int, guid string) error {
	/*	span := jaeger.GetSpan(ctx, "AddUserFavourites")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `insert into user_favourites ("name",json_array,guid,user_id)
	select uf."name", uf.json_array, get_uuid(), $2 from user_favourites uf where uf.guid = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "guid": guid}, "AddUserFavourites")

	_, err := db.GetClient().ExecContext(ctx, query, guid, userId)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "guid": guid}, "AddUserFavourites error", err)
		return err
	}

	return nil
}

func (db *Postgres) DeleteUserFavourites(ctx context.Context, userId int, guid string) error {
	/*	span := jaeger.GetSpan(ctx, "DeleteUserFavourites")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `delete from user_favourites where user_id = $2 and guid = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "guid": guid}, "DeleteUserFavourites")

	_, err := db.GetClient().ExecContext(ctx, query, guid, userId)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "guid": guid}, "DeleteUserFavourites error", err)
		return err
	}

	return nil
}

func (db *Postgres) RenameUserFavourites(ctx context.Context, userId int, request model.RenameUserFavouritesRequest) error {
	/*	span := jaeger.GetSpan(ctx, "RenameUserFavourites")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `update user_favourites set name = $3 where user_id = $2 and guid = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "request": request}, "RenameUserFavourites")

	_, err := db.GetClient().ExecContext(ctx, query, request.Guid, userId, request.Name)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "request": request}, "RenameUserFavourites error", err)
		return err
	}

	return nil
}

func (db *Postgres) ChangeUserFavouritesItems(ctx context.Context, userId int, request model.ChangeUserFavouritesItemsRequest) error {
	/*	span := jaeger.GetSpan(ctx, "ChangeUserFavouritesItems")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	args := []interface{}{
		request.Url,
		userId,
	}

	query := ``
	if request.Guid != nil {
		if request.ForAdd {
			query = `update user_favourites set json_array = coalesce(json_array::jsonb,'[]'::jsonb) || jsonb_build_object('name', $3, 'url', $1) where user_id = $2 and guid = $4`
			args = append(args, request.PageName)
			args = append(args, request.Guid)
		} else {
			query = `update user_favourites set json_array = (
				select jsonb_agg(jsonb_build_object('url', url, 'name', name))
					from jsonb_to_recordset((select json_array::jsonb from user_favourites uf where user_id = $2 and guid = $3)) as favourites(url varchar, name varchar)
					where url != $1
				)
				where user_id = $2 and guid = $3`
			args = append(args, request.Guid)
		}
	} else {
		query = `with i as (select get_uuid() as id)
		insert into user_favourites ("name",json_array,guid,user_id) values
		((select id from i),'[]'::jsonb || jsonb_build_object('name', $3, 'url', $1),(select id from i),$2)`
		args = append(args, request.PageName)
	}

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "request": request}, "ChangeUserFavouritesItems")

	_, err := db.GetClient().ExecContext(ctx, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "request": request}, "ChangeUserFavouritesItems error", err)
		return err
	}

	return nil
}

func (db *Postgres) Logout(ctx context.Context, userId int) error {
	/*	span := jaeger.GetSpan(ctx, "Logout")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `update "user" set token = null where id = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId}, "Logout")

	_, err := db.GetClient().ExecContext(ctx, query, userId)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId}, "Logout error", err)
		return err
	}

	return nil
}
