# Скриншоты для отчёта

Сюда кладутся скрины после локальной проверки. Имена файлов совпадают со ссылками в `Project_template.md`.

| Файл | Когда сделать |
|------|----------------|
| `postman-local.png` | `cd tests/postman && npm run test:local` (сервисы из docker compose) |
| `kafka-ui.png` | http://localhost:8090 — топики movie/user/payment-events |
| `api-movies-ingress.png` | `curl https://cinemaabyss.example.com/api/movies` в Kubernetes |
| `events-service-logs.png` | `kubectl -n cinemaabyss logs deploy/events-service` после `npm run test:kubernetes` |
| `helm-pods.png` | `kubectl get pods -n cinemaabyss` после `helm install` |
| `helm-api-movies.png` | ответ API через ingress после helm |
| `circuit-breaker-fortio.png` | вывод fortio load при тесте Istio circuit breaker |

Если PNG ещё нет — сначала прогоните команды из README и `Project_template.md`, затем сохраните скрин с указанным именем.
