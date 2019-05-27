from flask import jsonify, request

from ..core.utils import github
from . import handlers


def _unsupported_event_error():
    return "Unknown or unsupported event", 400


@github.check_request
def dispatch(payload: dict):
    event_name = request.headers.get("X-Github-Event", None)

    if event_name == "pull_request":
        handler = handle_pull_request
    elif event_name == "ping":
        handler = handle_ping
    else:
        return _unsupported_event_error()

    return handler(payload=payload)


def handle_pull_request(*, payload: dict):
    action_name = payload.get("action", None)

    if action_name == "synchronize":
        handler = handlers.handle_synchronize_event
    else:
        return _unsupported_event_error()

    return handler(payload=payload)


def handle_ping(**_kwargs):
    return jsonify({"status": 200})
