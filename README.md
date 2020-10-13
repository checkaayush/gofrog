<div align="center">
  <p>
    <img src="https://user-images.githubusercontent.com/4137581/95133604-3072c100-077f-11eb-9c7d-5d8951cf0454.png" height="130px"/>
  </p>
  
  <h1>gofrog</h1>
</div>

## Introduction

gofrog is a service to fetch most downloaded artifacts for given repository on an Artifactory instance. It uses the [Artifactory REST API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API).

- **Robust CI/CD pipeline**: Linting, unit tests, builds and deploys are done using GitHub actions
- **Deployment**: [gofrog.herokuapp.com](gofrog.herokuapp.com) (Temporary credentials: Username=gofrog Password=gofrog)
- **Approach**: Fetching the top 2 most popular artifacts for a given repository has been implemented using a max-heap so that the solution is extensible to fetching top 'n' most popular artifacts. It also utilizes Go's excellent concurrency primitives to speed up processing artifact details for given repository.

### API Specification

- GET `/v1/health` Health check to indicate API health
- GET `/v1/mostPopular?repo=jcenter-cache&count=2` Fetches top 2 most downloaded artifacts in repo `jcenter-cache`

Sample JSON Response:
```json
{
  "results": [
    {
      "path": "jcenter-cache/org/hamcrest/hamcrest-parent/1.3/hamcrest-parent-1.3.pom",
      "downloadLink": "http://<IP_ADDRESS>:80/artifactory/jcenter-cache/org/hamcrest/hamcrest-parent/1.3/hamcrest-parent-1.3.pom",
      "downloads": 30
    },
    {
      "path": "jcenter-cache/org/apache/struts/struts2-core/2.3.14/struts2-core-2.3.14.pom",
      "downloadLink": "http://<IP_ADDRESS>:80/artifactory/jcenter-cache/org/apache/struts/struts2-core/2.3.14/struts2-core-2.3.14.pom",
      "downloads": 26
    }
  ],
  "error": ""
}
```

## Development

> Pre-requisites: Go 1.15+

1. Clone the repository locally.

2. From repository root, run:

```bash
make start
```

3. API will be up and running at http://localhost:5000.

## Testing

From repository root, run:

```bash
make test
```
