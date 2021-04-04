// environments/prom-grafana/dev
(import "ksonnet-util/kausal.libsonnet") +
(import "todo/todo.libsonnet") +
{
    _config+:: {
        todo: {
            name: "todo",
            image: "gcr.io/sandbox-project-tc/todo:latest",
            replicaCount: 3,
            dbUsername: "ZGVtbw==",
            dbPassword: "MTIzNA==",
            port: 8000
        },
    },
}
