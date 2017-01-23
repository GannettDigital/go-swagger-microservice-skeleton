FROM paas-docker-artifactory.gannettdigital.com/paas-centos7-base:latest
ARG BRANCH
ADD {{ RepoName }}-$BRANCH /usr/local/bin/{{ RepoName }}
RUN chmod a+rx /usr/local/bin/{{ RepoName }}
CMD ["{{ RepoName }}", "--host", "0.0.0.0", "--port", "8080"]
