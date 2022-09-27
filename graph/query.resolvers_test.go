package graph_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_Tasks(t *testing.T) {
	srv, client := NewGqlServer()
	q := `query ListTasks {
			  tasks(input: {
				title: "title"
				status: ""
			  }) {
				title
			  }
			}
		`
	initDBForTests(context.Background(), t, client)
	ctx := context.Background()
	require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)

	var resp map[string]interface{}
	srv.MustPost(q, &resp)
	_, exist := resp["tasks"]
	assert.True(t, exist)
	assert.Equal(t, "title", resp["tasks"].([]interface{})[0].(map[string]interface{})["title"])
}
