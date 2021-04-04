(import "ksonnet-util/kausal.libsonnet") +
{
    _config+:: {
        todo: {
            name: "todo",
            image: "gcr.io/sandbox-project-tc/todo:latest",
            replicaCount: 1,
            dbUsername: "ZGVtbw==",
            dbPassword: "MTIzNA==",
            port: 8000
        },
    },
}
+
{
    // use locals to extract the parts we need
    local deploy = $.apps.v1.deployment,
    local container = $.core.v1.container,
    local port = $.core.v1.containerPort,
    local service = $.core.v1.service,
    
    todo: {
        deployment: deploy.new(name=$._config.todo.name, replicas=$._config.todo.replicaCount, containers=[
            container.new($._config.todo.name, $._config.todo.image)
            + 
            container.withPorts( // add ports to the container
                [port.new("http", $._config.todo.port)]
            )
            +
            container.withEnvFrom([
                { secretRef: { name: 'db-auth' } },
            ])
            +
            container.mixin.livenessProbe.httpGet.withPath("/") 
            + 
            container.mixin.livenessProbe.httpGet.withPort("http")
            +
            container.mixin.readinessProbe.httpGet.withPath("/") 
            + 
            container.mixin.readinessProbe.httpGet.withPort("http")
            ,
        ])
        +
        deploy.mixin.metadata.withLabels({
            'app.kubernetes.io/name': $._config.todo.name,
            'app.kubernetes.io/instance': $._config.todo.name,
            'app.kubernetes.io/managed-by': 'Tanka'
        })
        +
        deploy.mixin.metadata.withNamespace("dev")
        ,
        service: $.util.serviceFor(self.deployment) + service.mixin.metadata.withNamespace("dev"),
        secret: {
            apiVersion: 'v1',
            kind: 'Secret',
            metadata: {
                name: 'db-auth',
                labels: {
                    'app.kubernetes.io/name': $._config.todo.name,
                    'app.kubernetes.io/instance': $._config.todo.name,
                    'app.kubernetes.io/managed-by': 'Tanka'
                },
            },
            data: {
                POSTGRES_PASSWORD: $._config.todo.dbPassword,
                POSTGRES_USER: $._config.todo.dbUsername
            },
        },
    }
}

