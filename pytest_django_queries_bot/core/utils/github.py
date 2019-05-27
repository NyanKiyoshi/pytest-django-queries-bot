import hashlib
import hmac
from typing import Optional

import flask


def _match_signatures() -> bool:
    received_signature = flask.request.headers.get("HTTP_X_HUB_SIGNATURE")

    if received_signature is None:
        return False

    github_secret = flask.current_app.config.get(
        "GITHUB_SECRET_KEY"
    )  # type: Optional[bytes]

    if github_secret is None:
        return True

    body = bytes(flask.request.get_data(as_text=True), "utf-8")
    computed_signature = (
        "sha1=" + hmac.new(github_secret, body, hashlib.sha1).hexdigest()
    )

    return hmac.compare_digest(computed_signature, received_signature)


def check_request(func):
    def _check(*args, **kwargs):
        kwargs["payload"] = payload = flask.request.json

        if payload is None:
            return "Missing payload", 400

        if _match_signatures():
            return func(*args, **kwargs)

        return "Signatures don't match", 403

    return _check
