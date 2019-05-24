import json
from typing import Optional

import pytest
from flask import Response, url_for

from tests.conftest import get_headers


def get_payload(*, action=Optional[str]):
    payload = {}

    if action:
        payload["action"] = action

    return json.dumps(payload)


def test_without_payload(client):
    url = url_for("webhook")
    headers = get_headers(event="pull_request")
    headers["Content-Type"] = "oopsie"
    response = client.post(url, headers=headers)  # type: Response
    assert response.status_code == 400, next(response.response)
    assert next(response.response) == b"Missing payload."


@pytest.mark.parametrize("action", (None, "unk"))
def test_invalid_event(client, action):
    url = url_for("webhook")
    headers = get_headers(event="pull_request")
    payload = get_payload(action=action)
    response = client.post(url, headers=headers, data=payload)  # type: Response
    assert response.status_code == 400, next(response.response)
    assert next(response.response) == b"Unknown or unsupported event."


def test_valid_event(client):
    url = url_for("webhook")
    headers = get_headers(event="pull_request")
    payload = get_payload(action="synchronize")
    response = client.post(url, headers=headers, data=payload)  # type: Response
    assert response.status_code == 200, next(response.response)
