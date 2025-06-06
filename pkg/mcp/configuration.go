package mcp

import (
	"context"
	"fmt"
	"github.com/manusa/kubernetes-mcp-server/pkg/kubernetes"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) initConfiguration() []server.ServerTool {
	return []server.ServerTool{
		{mcp.NewTool("configuration_view",
			mcp.WithDescription("Get the current Kubernetes configuration content as a kubeconfig YAML"),
			mcp.WithBoolean("minified", mcp.Description("Return a minified version of the configuration. "+
				"If set to true, keeps only the current-context and the relevant pieces of the configuration for that context. "+
				"If set to false, all contexts, clusters, auth-infos, and users are returned in the configuration. "+
				"(Optional, default true)")),
		), configurationView},
	}
}

func configurationView(_ context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	minify := true
	minified := ctr.Params.Arguments["minified"]
	if _, ok := minified.(bool); ok {
		minify = minified.(bool)
	}
	ret, err := kubernetes.ConfigurationView(minify)
	if err != nil {
		err = fmt.Errorf("failed to get configuration: %v", err)
	}
	return NewTextResult(ret, err), nil
}
