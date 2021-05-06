from chalice import Blueprint, Cron
from chalicelib import _overrides
from chalicelib.utils import helper

app = Blueprint(__name__)
_overrides.chalice_app(app)
from chalicelib.ee import telemetry


# Run every day.
@app.schedule(Cron('0', '0', '?', '*', '*', '*'))
def telemetry_cron(event):
    telemetry.compute()
