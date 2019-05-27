import pytest
from flask import Response, url_for

from tests.conftest import get_headers


@pytest.mark.parametrize("payload", (None, "is this json?"))
def test_invalid_payload(client, payload):
    url = url_for("webhook")
    headers = get_headers(event="pull_request")
    response = client.post(url, headers=headers, data=payload)  # type: Response
    assert response.status_code == 400, next(response.response)


def test_invalid_http_method(client):
    url = url_for("webhook")
    response = client.get(url)  # type: Response
    assert response.status_code == 405, next(response.response)
    assert next(response.response) != b"Unknown or unsupported event."


@pytest.mark.parametrize("event", (None, "unk"))
def test_invalid_event(client, event, dummy_payload):
    url = url_for("webhook")
    headers = get_headers(event=event)
    response = client.post(url, headers=headers, data=dummy_payload)  # type: Response
    assert response.status_code == 400, next(response.response)
    assert next(response.response) == b"Unknown or unsupported event"


def test_ping_event(client, dummy_payload):
    url = url_for("webhook")
    headers = get_headers(event="ping")
    response = client.post(url, headers=headers, data=dummy_payload)  # type: Response
    assert response.status_code == 200, next(response.response)
    assert response.json == {"status": 200}
