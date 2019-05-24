from flask import Response, url_for


def test_get(client):
    url = url_for("webhook")
    response = client.get(url)  # type: Response
    assert response.status_code == 200
