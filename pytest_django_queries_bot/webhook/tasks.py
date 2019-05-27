from ..core.extensions import celery


@celery.task(name="webhooks.tasks.diff_generator")
def create_diff_for_repository():
    pass
