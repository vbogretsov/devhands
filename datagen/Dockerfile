FROM python:3.12-slim as build
RUN pip install -U pip
RUN pip install uv~=0.6
WORKDIR /src
ADD uv.lock .
ADD pyproject.toml .
RUN uv venv /opt
RUN . /opt/bin/activate \
&& uv sync --no-install-project --active

FROM python:3.12-slim
COPY --from=build /opt /opt
ENV PYTHONPATH /opt/lib/python3.12/site-packages
ENV PATH /opt/bin:${PATH}
RUN adduser --system --home /app app
WORKDIR /app
COPY . .
