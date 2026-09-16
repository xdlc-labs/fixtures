# Backlog

## TestBillingUpstream needs a real billing upstream

CI fails because `TestBillingUpstream` dials `http://billing.internal.example:9/ready`, and that host does not resolve in GitHub Actions (or any public runner).

This cannot be fixed from this repo alone. The check is a live readiness probe against an internal billing service that is not part of this codebase, not mocked here, and not reachable from CI. Making the suite green would mean deleting the test or faking a 200, which would hide the real dependency rather than satisfy it.

**Unblock:** run this check only where `billing.internal.example` is available (private network / staging), or provide a documented test double / contract that the billing team owns. Until then, leave the assertion in place and keep CI red for this gate.
