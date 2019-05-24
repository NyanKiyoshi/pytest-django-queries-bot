import pytest
from flask import Response, url_for

from tests.conftest import get_headers


def test_invalid_http_method(client):
    url = url_for("webhook")
    response = client.get(url)  # type: Response
    assert response.status_code == 405
    assert next(response.response) != b"Unknown or unsupported event."


@pytest.mark.parametrize("event", (None, "unk"))
def test_invalid_event(client, event):
    url = url_for("webhook")
    headers = get_headers(event=event)
    response = client.post(url, headers=headers)  # type: Response
    assert response.status_code == 400
    assert next(response.response) == b"Unknown or unsupported event."


def test_ping_event(client):
    url = url_for("webhook")
    headers = get_headers(event="ping")
    response = client.post(url, headers=headers)  # type: Response
    assert response.status_code == 200
    assert response.json == {"status": 200}
