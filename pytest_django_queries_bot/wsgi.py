from pytest_django_queries_bot.application import celery, create_app

application = create_app()

__all__ = ["application", "celery"]
