# Backlog

## TestBillingUpstream needs a live billing host

CI fails on `TestBillingUpstream` (`src/billing_test.go`): it GETs
`http://billing.internal.example:9/ready` and fatals on dial/DNS errors.

`billing.internal.example` does not resolve on GitHub Actions (or any typical
laptop). Nothing in this repo can make that host answer without deleting the
test, stubbing HTTP, or pointing at a fake URL. Those would green CI by lying
about the upstream the test is meant to assert.

**Blocked on:** a real billing service (or a shared test double) reachable at
that hostname/port in CI, owned outside this fixtures app. Until then, keep CI
red and do not “fix” the assertion away.
