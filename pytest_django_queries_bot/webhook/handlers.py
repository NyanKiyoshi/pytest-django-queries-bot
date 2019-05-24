"""See https://developer.github.com/v3/issues/events/ for references."""
from flask import jsonify, request

from . import actions as pr_actions_handlers


def handle_pull_request_event():
    payload = request.json

    if payload is None:
        return "Missing payload.", 400

    if "action" in payload:
        handler = getattr(
            pr_actions_handlers,
            f"handle_{payload['action']}_event",
            handle_unsupported_event,
        )
    else:
        handler = handle_unsupported_event

    return handler(payload)


def handle_ping_event():
    return jsonify({"status": 200})


def handle_unsupported_event(*_):
    return "Unknown or unsupported event.", 400


DEFAULT_EVENT_HANDLER = handle_unsupported_event
