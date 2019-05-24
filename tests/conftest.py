from os.path import dirname, join, realpath
from unittest import mock

import pytest
from jinja2 import StrictUndefined

from pytest_django_queries_bot.application import create_app


@pytest.fixture(scope="session")
def app():
    config_path = join(realpath(dirname(__file__)), "settings.py")
    with mock.patch(
        "pytest_django_queries_bot.application.get_config_path"
    ) as mocked_cfg_path:
        mocked_cfg_path.return_value = config_path
        flask_app = create_app()
        mocked_cfg_path.assert_called_once_with()

    # make jinja raise an exception on undefined variables
    flask_app.jinja_env.undefined = StrictUndefined()
    return flask_app


@pytest.fixture
def pushed_app(app):
    app_context = app.app_context()
    app_context.push()

    yield app
    app_context.pop()


@pytest.fixture
def client(pushed_app):
    return pushed_app.test_client()
