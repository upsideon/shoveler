## Shoveler

When it snows, there are people who have difficulty or are unable to shovel the walkway at their place of residence. Some have family members, friends, or contractors to help them with this task. For those who don't have these resources, Shoveler seeks to be a solution.

The goals of this project, which is still in its infancy, are as follows:

- If you are in need of snow shoveling assistance, there is a dedicated venue in which to make that need known and fulfilled.
- If you have a desire to help others in your community, providing snow shoveling assistance is a clear and present option.

There is no business model here. If this becomes something more than code on a Git repository, the intention is to run the service as a non-profit via donations. This project was inspired by the [Slim.AI Hackathon](https://slim-ai-hackathon.devpost.com/).

### Development

All services are containerized using Docker. If you haven't installed Docker already, please follow the instructions [here](https://docs.docker.com/engine/install/) before proceeding.

If it is your first time running Shoveler, execute the following commands from the top-level directory to bring it up:

```
docker-compose up -d cassandra
docker exec -i shoveler_cassandra_1 cqlsh < database/migrations.cql
docker-compose up -d shoveler-frontend shoveler-backend
```

On subsequent runs, as long as you don't remove the Cassandra container, you can execute the following commands to bring up Shoveler:

```
docker-compose up -d cassandra && sleep 60 && docker-compose up --build
```

The sleep command gives the Cassandra database enough time to come up before the Shoveler backend attempts to access it. In the future, either intelligent retries by the backend or health checks to support `depends_on` in `docker-compose.yml` will be added to address this problem.

The web page can be accessed by going to your browser and navigating to [http://localhost:3000](http://localhost:3000). To access the backend service directly, make HTTP requests to [http://localhost:8080](http://localhost:8080). If these ports are already in use on your system, feel free to update the port mapping in `docker-compose.yml` to meet your needs.

When you're ready to stop Shoveler, execute the following command from the top-level directory:

```
docker-compose stop
```

### Deploying Container Images

To deploy container images to DockerHub for the frontend and backend services, first add the following environment variable export to the configuration script for your favorite shell. An example for Bash is provided below:

```
echo "export SHOVELER_DOCKERHUB_REPO=<dockerhub-repo>" >> ~/.bashrc
```

Then execute `dockerhub-push.sh` build the images and push them up to DockerHub.

### Security Assessment

#### Frontend

The frontend was written using JavaScript via the React.js framework. A vulnerability scan of the `shoveler-frontend` image reveals the following security vulnerabilities:

|Risk Level|Summary|Occurrences|Fix availability|Vulnerability|
|----------|-------|-----------|----------------|-------------|
| Critical | This affects all versions of package json-pointer. A type confusion vulnerability can lead to a bypass of CVE-2020-7709 when the pointer components are arrays. | 1 | Unknown | CVE-2021-23820 |
| High | nth-check is vulnerable to Inefficient Regular Expression Complexity | 2 | Fixed | CVE-2021-3803 |
| High | websocket-extensions ruby module prior to 0.1.5 allows Denial of Service (DoS) via Regex Backtracking. The extension parser may take quadratic time when parsing a header containing an unclosed string parameter value whose content is a repeating two-byte sequence of a backslash and some other character. This could be abused by an attacker to conduct Regex Denial Of Service (ReDoS) on a single-threaded server by providing a malicious payload with the Sec-WebSocket-Extensions header. | 1 | Unknown | CVE-2020-7663|
| High | An issue was discovered in the cookie crate before 0.7.6 for Rust. Large integers in the Max-Age of a cookie cause a panic. | 1 | Unknown | CVE-2017-18589|
| Medium | JSDom improperly allows the loading of local resources, which allows for local files to be manipulated by a malicious web page when script execution is enabled. | 1 | Unknown | CVE-2021-20066 |
| Medium | The quick login feature in Slash Slashcode does not redirect the user to an alternate URL when the wrong password is provided, which makes it easier for remote web sites to guess the proper passwords by reading the username and password from the Referrer URL. | Unknown | CVE-2002-1647 |
| Low | Cross-site scripting (XSS) vulnerability in the 3 slide gallery in the Fresh theme before 7.x-1.4 for Drupal allows remote authenticated users with the administer themes permission to inject arbitrary web script or HTML via unspecified vectors. | 1 | Unknown | CVE-2013-1779|

Additional details on these vulnerabilities can be obtained by viewing the full JSON output in `security/frontend_vulnerabilities.json`. The Slim.AI platform allows these files to be exported, which is useful, but it would be nice to have the ability to import such a JSON file for viewing as well.

For the majority of these vulnerabilities, it is unknown whether a fix is available according to the Grype and Trivy scanners. The `nth-check` is the only vulnerability that has a patch. This fact is confirmed by running `npm audit`. Running `npm audit fix --force` does not fix this problem for us, however. This is due to the incompatibility of the dependencies with such an upgrade.

Fortunately, we can avoid these vulnerabilities in the meantime by ensuring that we avoid using the parts of the libraries that utilize them.

#### Backend


The backend was written in Go using the `gin-gonic` web framework. A vulnerability scan of the `shoveler-backend` image reveals the following security vulnerabilities:

|Risk Level|Summary|Occurrences|Fix availability|Vulnerability|
|----------|-------|-----------|----------------|-------------|
| High | golang: golang.org/x/text/language: ParseAcceptLanguage takes a long time to parse complex tags | 1 | Fixed | CVE-2022-32149	|
| High | golang: net/http: handle server errors after sending GOAWAY | 1 | Fixed | CVE-2022-27664	|
| High | golang: crash in a golang.org/x/crypto/ssh server | 1 | Fixed | CVE-2022-27191 |
| High | golang: net/http: limit growth of header canonicalization cache | 1 | Fixed | CVE-2021-44716	|
| High | golang.org/x/crypto: empty plaintext packet causes panic | 1	| Fixed | CVE-2021-43565 |
| High | golang: out-of-bounds read in golang.org/x/text/language leads to DoS | 1 | Fixed | CVE-2021-38561	|
| High | golang: x/net/html: infinite loop in ParseFragment | 1	| Fixed | CVE-2021-33194 |
| High | protobuf allows remote authenticated attackers to cause a heap-based buffer overflow. | 1 | Unknown | CVE-2015-5237 |
| Medium | golang: syscall: faccessat checks wrong group | 1 | Fixed | CVE-2022-29526	|
| Medium | golang: net/http: panic in ReadRequest and ReadResponse when reading a very large header | 1	| Fixed | CVE-2021-31525	|
| Medium | Nullptr dereference when a null char is present in a proto symbol. The symbol is parsed incorrectly, lead ... ce the symbol is incorrectly parsed, the file is nullptr. We recommend upgrading to version 3.15.0 or greater. | 1	| Unknown | CVE-2021-22570 |

Additional details on these vulnerabilities can be obtained by viewing the full JSON output in `security/backend_vulnerabilities.json`.

Many of these vulnerabilities are present in the standard Go library and marked as fixed even though the Docker image is on the latest tagged version of 1.19.3. This is rather odd and makes it so that we can't resolve them simply by upgrading the version of Go used by the image at this time. The solution is once again to be knowledgeable of what portions of these libraries are vulnerable and avoid using them.
