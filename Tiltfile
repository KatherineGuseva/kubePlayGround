load('ext://restart_process', 'docker_build_with_restart')

docker_build_with_restart('merge', '.', 
    dockerfile='Dockerfile.dev',
    entrypoint='/go-merge merge',
    live_update=[
        run('--chown=1001:0 --from=build /app/go-merge .'),
    ]
)
k8s_yaml('kubePayGround/merge.yaml')
