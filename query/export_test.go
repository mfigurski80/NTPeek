package query

/// This file contains test helpers that are not part of the public API.
/// They are exported for ease of mocking: they SHOULD NOT BE TESTED DIRECTLY.

func SET_HTTP_CLIENT(c httpClientType) {
	httpClient = c
}
