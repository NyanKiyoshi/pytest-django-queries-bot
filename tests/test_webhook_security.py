import hashlib
import hmac

from flask import Response, url_for

from tests.conftest import get_headers


def test_signature_is_required_if_set(client, config, dummy_payload):
    url = url_for("webhook")
    config["GITHUB_SECRET_KEY"] = b"dummy"
    headers = get_headers(event="ping")
    response = client.post(url, headers=headers, data=dummy_payload)  # type: Response
    assert response.status_code == 403, next(response.response)
    assert next(response.response) == b"Signatures don't match"

    signature = hmac.new(
        b"dummy", bytes(dummy_payload, "utf-8"), hashlib.sha1
    ).hexdigest()
    headers = get_headers(event="ping", signature="sha1=" + signature)
    response = client.post(url, headers=headers, data=dummy_payload)  # type: Response
    assert response.status_code == 200, next(response.response)
    assert b": 200" in next(response.response)


def test_signature_is_required_if_set_missing_header_is_handled(
    client, config, dummy_payload
):
    url = url_for("webhook")
    config["GITHUB_SECRET_KEY"] = b"blah"
    headers = get_headers(event="ping")
    headers.pop("HTTP_X_HUB_SIGNATURE")
    response = client.post(url, headers=headers, data=dummy_payload)  # type: Response
    assert response.status_code == 403, next(response.response)
    assert next(response.response) == b"Signatures don't match"
