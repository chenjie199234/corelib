package mrpc

import (
	"context"
)

type inmetadatakey struct{}

//in
func GetAllInMetadata(ctx context.Context) map[string]string {
	value := ctx.Value(inmetadatakey{})
	if value == nil {
		return nil
	}
	result, ok := value.(map[string]string)
	if !ok {
		return nil
	}
	return result
}
func GetInMetadata(ctx context.Context, key string) string {
	value := ctx.Value(inmetadatakey{})
	if value == nil {
		return ""
	}
	result, ok := value.(map[string]string)
	if !ok {
		return ""
	}
	return result[key]
}
func SetInMetadata(ctx context.Context, key, value string) context.Context {
	tempresult := ctx.Value(inmetadatakey{})
	if tempresult == nil {
		return context.WithValue(ctx, inmetadatakey{}, map[string]string{key: value})
	}
	result := tempresult.(map[string]string)
	result[key] = value
	ctx = context.WithValue(ctx, inmetadatakey{}, result)
	return ctx
}
func SetAllInMetadata(ctx context.Context, data map[string]string) context.Context {
	tempresult := ctx.Value(inmetadatakey{})
	if tempresult == nil {
		return context.WithValue(ctx, inmetadatakey{}, data)
	}
	result := tempresult.(map[string]string)
	for k, v := range data {
		result[k] = v
	}
	ctx = context.WithValue(ctx, inmetadatakey{}, result)
	return ctx
}

type outmetadatakey struct{}

//out
func GetAllOutMetadata(ctx context.Context) map[string]string {
	value := ctx.Value(outmetadatakey{})
	if value == nil {
		return nil
	}
	result, ok := value.(map[string]string)
	if !ok {
		return nil
	}
	return result
}
func GetOutMetadata(ctx context.Context, key string) string {
	value := ctx.Value(outmetadatakey{})
	if value == nil {
		return ""
	}
	result, ok := value.(map[string]string)
	if !ok {
		return ""
	}
	return result[key]
}
func SetOutMetadata(ctx context.Context, key, value string) context.Context {
	tempresult := ctx.Value(outmetadatakey{})
	if tempresult == nil {
		return context.WithValue(ctx, outmetadatakey{}, map[string]string{key: value})
	}
	result := tempresult.(map[string]string)
	result[key] = value
	ctx = context.WithValue(ctx, outmetadatakey{}, result)
	return ctx
}
func SetAllOutMetadata(ctx context.Context, data map[string]string) context.Context {
	tempresult := ctx.Value(outmetadatakey{})
	if tempresult == nil {
		return context.WithValue(ctx, outmetadatakey{}, data)
	}
	result := tempresult.(map[string]string)
	for k, v := range data {
		result[k] = v
	}
	ctx = context.WithValue(ctx, outmetadatakey{}, result)
	return ctx
}
