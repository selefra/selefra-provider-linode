package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeKubernetesClusterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeKubernetesClusterGenerator{}

func (x *TableLinodeKubernetesClusterGenerator) GetTableName() string {
	return "linode_kubernetes_cluster"
}

func (x *TableLinodeKubernetesClusterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeKubernetesClusterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeKubernetesClusterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeKubernetesClusterGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListLKEClusters(ctx, nil)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range items {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getKubeConfig(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	cluster := result.(linodego.LKECluster)
	conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
	if err != nil {

		return nil, err
	}
	item, err := conn.GetLKEClusterKubeconfig(ctx, cluster.ID)
	if err != nil {

		return nil, err
	}
	return item, err
}
func listKubernetesClusterAPIEndpoints(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	cluster := result.(linodego.LKECluster)
	conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
	if err != nil {

		return nil, err
	}
	items, err := conn.ListLKEClusterAPIEndpoints(ctx, cluster.ID, nil)
	if err != nil {

		return nil, err
	}
	endpoints := []string{}
	for _, i := range items {
		endpoints = append(endpoints, i.Endpoint)
	}
	return endpoints, err
}
func listKubernetesClusterPools(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	cluster := result.(linodego.LKECluster)
	conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
	if err != nil {

		return nil, err
	}
	items, err := conn.ListLKENodePools(ctx, cluster.ID, nil)
	if err != nil {

		return nil, err
	}
	return items, err
}

func (x *TableLinodeKubernetesClusterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeKubernetesClusterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags applied to the Kubernetes cluster as a map.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("List of Tags applied to the Kubernetes cluster.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("This Kubernetes cluster’s unique label for display purposes only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_endpoints").ColumnType(schema.ColumnTypeJSON).Description("API endpoints for the cluster.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listKubernetesClusterAPIEndpoints(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("This Kubernetes cluster’s location.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kubeconfig").ColumnType(schema.ColumnTypeString).Description("Kube config for the cluster.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getKubeConfig(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("KubeConfig")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pools").ColumnType(schema.ColumnTypeJSON).Description("Pools for the cluster.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listKubernetesClusterPools(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("When this Kubernetes cluster was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("This Kubernetes cluster’s unique ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("When this Kubernetes cluster was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("k8s_version").ColumnType(schema.ColumnTypeString).Description("The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>, and the latest supported patch version will be deployed.").
			Extractor(column_value_extractor.StructSelector("K8sVersion")).Build(),
	}
}

func (x *TableLinodeKubernetesClusterGenerator) GetSubTables() []*schema.Table {
	return nil
}
