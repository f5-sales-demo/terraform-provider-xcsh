---
page_title: "xcsh_healthcheck examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck examples."
---

# xcsh_healthcheck examples

<a id="canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- Examples

<a id="canonical-2022022332323212-3130212231000313-3201300210131003-0203220222002221-2001312112212311-2101220201023322-2220322010313101-3223031110322013"></a>

### Complete configurations for `xcsh_healthcheck`

- [All attributes](resources--healthcheck--examples--group-001.md#canonical-2310113000112303-1100022330022013-1320103331202232-0000111010322012-1002230021020103-1223330100120030-1012033223013011-2101122112233010): valid configuration.

- [Http headers remove](resources--healthcheck--examples--group-001.md#canonical-0302300320021011-0003132121331002-3013303330130120-0123320022301312-1323133301233312-1010322312322013-0033133013113311-3101222030021211): valid configuration.

- [Http health check](resources--healthcheck--examples--group-001.md#canonical-0112102133233221-2002233010320103-1032320003320102-3313223330322100-2130230101013030-1031011312021212-1021223231113210-0331233020020330): valid configuration.

- [Http http2](resources--healthcheck--examples--group-001.md#canonical-1310130002020132-0022010320210002-3021310133232030-0222300211223012-3231010102122112-0322031231013030-1321313213113031-2020333321313202): valid configuration.

- [Http origin server name](resources--healthcheck--examples--group-001.md#canonical-3302000302103322-3133323211121002-2013031220203013-3203002312133122-2020010112003311-3130323123110230-3023002320103222-3122012102133021): valid configuration.

- [Http status codes](resources--healthcheck--examples--group-001.md#canonical-1202313013211210-0230021022022211-1331222022313002-3231031032013222-0332322320203331-3020122022001111-0011233202220330-2131233012232112): valid configuration.

- [Http with path](resources--healthcheck--examples--group-001.md#canonical-1202320233200322-2131113312123320-0321302112011200-1122112032311122-2000203223302000-0232222333320202-0110303000333212-3301300203311231): valid configuration.

- [Resource](resources--healthcheck--examples--group-001.md#canonical-1121001010013121-0113111201111232-2322101113010012-3122320331121322-1332200131001213-2131212311321203-3222023231213003-1112001123100303): valid configuration.

- [Thresholds](resources--healthcheck--examples--group-001.md#canonical-1022121031301321-3032210032102310-2003203223112330-3320322300010310-0101222300012320-0231110320011312-3231113132211233-1102111121110131): valid configuration.

- [Udp icmp](resources--healthcheck--examples--group-001.md#canonical-0202330033310230-2122330132020231-2232010012202201-3210330020023002-1131302202211123-3103323223131111-0000203121030110-0033112011013330): valid configuration.

- [With annotations](resources--healthcheck--examples--group-001.md#canonical-2000033331200313-1301221012330111-0001213033031230-1013302321322233-1333001013232020-1203032331312331-2101230233031211-1003332302320131): valid configuration.

- [With description](resources--healthcheck--examples--group-001.md#canonical-2121020122222332-0011213211310322-3012333301201312-2101101333210123-3320220133122120-1213300012111102-3113130021202212-1203100110103300): valid configuration.

- [With jitter](resources--healthcheck--examples--group-001.md#canonical-1213303100211112-0310221121330321-0300210312002303-1112302300120312-0120211313122313-0120130232221320-0220102330331211-1000013020212032): valid configuration.

- [With labels](resources--healthcheck--examples--group-001.md#canonical-3012001213313201-3012312021301211-2112132102130210-1123133113132122-1010220023210130-2121202221031221-3222003032322003-3021300311323332): valid configuration.

<a id="canonical-2310113000112303-1100022330022013-1320103331202232-0000111010322012-1002230021020103-1223330100120030-1012033223013011-2101122112233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/all-attributes.tf`; digest `sha256:7b4a2a1e5d67895b93aa7dcead7285790016b97f45b00fc3456985658bd0bba4`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  tcp_health_check {}
}
```

<a id="canonical-0302300320021011-0003132121331002-3013303330130120-0123320022301312-1323133301233312-1010322312322013-0033133013113311-3101222030021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http headers remove example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http headers remove

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-headers-remove.tf`; digest `sha256:f6be554fcc3598ae81c69deefdb57d62d8b7fb9365a310046f4973eeaeb364af`.

```terraform
# HttpHeadersRemove — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path                      = "example-value"
    host_header               = "example.com"
    request_headers_to_remove = ["X-Custom-Header", "X-Debug"]
  }
}
```

<a id="canonical-0112102133233221-2002233010320103-1032320003320102-3313223330322100-2130230101013030-1031011312021212-1021223231113210-0331233020020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http health check example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http health check

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-health-check.tf`; digest `sha256:d687dc5e245ad34528fd20feda8e0150dd05ca683b1bddfd84c19f770f54453d`.

```terraform
# HttpHealthCheck — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "/health"
    host_header = "example.com"
  }
}
```

<a id="canonical-1310130002020132-0022010320210002-3021310133232030-0222300211223012-3231010102122112-0322031231013030-1321313213113031-2020333321313202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http http2 example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http http2

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-http2.tf`; digest `sha256:7586adcf4b2cb6423fa64e91dd0205c03d4130e45c575432d8fff8d1cf0d1e8f`.

```terraform
# HttpHttp2 — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "example-value"
    host_header = "example.com"
    use_http2   = true
  }
}
```

<a id="canonical-3302000302103322-3133323211121002-2013031220203013-3203002312133122-2020010112003311-3130323123110230-3023002320103222-3122012102133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http origin server name example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http origin server name

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-origin-server-name.tf`; digest `sha256:afde24e907d59c17ebd5e445c7c324b1ec4a0b83e3d3650974943970ea9a430a`.

```terraform
# HttpOriginServerName — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path                   = "example-value"
    use_origin_server_name = {}
  }
}
```

<a id="canonical-1202313013211210-0230021022022211-1331222022313002-3231031032013222-0332322320203331-3020122022001111-0011233202220330-2131233012232112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http status codes example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http status codes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-status-codes.tf`; digest `sha256:22e40413d5391679e83e76d2d3211aa82db4b3b9529e4270b634ac5f37238cfb`.

```terraform
# HttpStatusCodes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path                  = "example-value"
    host_header           = "example.com"
    expected_status_codes = ["200", "201", "204"]
  }
}
```

<a id="canonical-1202320233200322-2131113312123320-0321302112011200-1122112032311122-2000203223302000-0232222333320202-0110303000333212-3301300203311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Http with path example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Http with path

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-with-path.tf`; digest `sha256:46a31aa97f576bd69e9ed469c71f60122351ddb6f19015d16952ce78f199eb8e`.

```terraform
# HttpWithPath — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  http_health_check {
    path        = "example-value"
    host_header = "example.com"
  }
}
```

<a id="canonical-1121001010013121-0113111201111232-2322101113010012-3122320331121322-1332200131001213-2131212311321203-3222023231213003-1112001123100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/resource.tf`; digest `sha256:3cd7562f625a6457911202f74161bf5f6a0eec910fe3f71c8a9f0c0e1ea3d639`.

```terraform
# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Healthcheck configuration
resource "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  healthy_threshold   = 1
  interval            = 1
  timeout             = 1
  unhealthy_threshold = 1
}
```

<a id="canonical-1022121031301321-3032210032102310-2003203223112330-3320322300010310-0101222300012320-0231110320011312-3231113132211233-1102111121110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Thresholds example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Thresholds

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/thresholds.tf`; digest `sha256:eb7ffc7c4c13b7cf863658813f0b685ce60bc365c197904650a919fd4f6c683f`.

```terraform
# Thresholds — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 3
  unhealthy_threshold = 2
  timeout             = 5
  interval            = 15

  tcp_health_check {}
}
```

<a id="canonical-0202330033310230-2122330132020231-2232010012202201-3210330020023002-1131302202211123-3103323223131111-0000203121030110-0033112011013330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Udp icmp example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- Udp icmp

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/udp-icmp.tf`; digest `sha256:a9f4239096a1306a82a2b1a05588ca6a7a9bf652a5283a4195352053169e6fb7`.

```terraform
# UdpIcmp — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  udp_icmp_health_check = {}
}
```

<a id="canonical-2000033331200313-1301221012330111-0001213033031230-1013302321322233-1333001013232020-1203032331312331-2101230233031211-1003332302320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With annotations example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-annotations.tf`; digest `sha256:f3193cdc60afd582627e020475994d811ba4d8f1838c6f8e15c607099f78229e`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  annotations = {
    key1 = "example-value"
    key2 = "example-description"
  }

  tcp_health_check {}
}
```

<a id="canonical-2121020122222332-0011213211310322-3012333301201312-2101101333210123-3320220133122120-1213300012111102-3113130021202212-1203100110103300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With description example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-description.tf`; digest `sha256:645f65af9f47887d2bb355f22d93ee442149a4f36d831425d41955bb04adbff3`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name        = "example"
  namespace   = "system"
  description = "example-value"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  tcp_health_check {}
}
```

<a id="canonical-1213303100211112-0310221121330321-0300210312002303-1112302300120312-0120211313122313-0120130232221320-0220102330331211-1000013020212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With jitter example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- With jitter

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-jitter.tf`; digest `sha256:b066d263997260c6a34ca816fcfdef53d425a0fd6680007fb1a8a2d0ba72d4fe`.

```terraform
# WithJitter — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5
  jitter_percent      = 30

  tcp_health_check {}
}
```

<a id="canonical-3012001213313201-3012312021301211-2112132102130210-1123133113132122-1010220023210130-2121202221031221-3222003032322003-3021300311323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md#canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100)
- [Examples](resources--healthcheck--examples--group-001.md#canonical-2313303123101000-3213111110100003-0220003232111323-1332220121322330-3030130111203200-0302231311033103-3223333012201330-0011221311203111)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-labels.tf`; digest `sha256:9d101becb63d5e10a0b0292248bcfa96ec4395ca9743eb0d33275bb701472643`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "example-value"
    managed_by  = "example-description"
  }

  tcp_health_check {}
}
```
