Библиотека используется для автоматического версионирования GO-приложения с использованием переменных gitlab

В dockerfile надо изменить строку билда на такую:
```
RUN go build -ldflags "-X github.com/PCManiac/compile_vars.version=${CI_PIPELINE_IID} -X github.com/PCManiac/compile_vars.build_time=`date +%FT%T%z`" -o server ./...
```

Ну и в .gitlab-ci.yml собирать докер с подстановкой переменных
```
- docker build --build-arg CI_PIPELINE_IID="$CI_PIPELINE_IID" --build-arg GITLAB_TOKEN="$CI_JOB_TOKEN" --pull -t "$CI_REGISTRY_IMAGE" .
```

При инициализации модуля в log будет добавлена строка с номером версии и временем билда. При необходимости эти значения также можно получить с помощью функций GetVersion() и GetBuildTime()