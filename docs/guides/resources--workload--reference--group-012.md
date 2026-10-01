---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1323022002023322-2122300023210210-0000113002232013-3031011301022022-3311233110230012-1121032300221110-2001120011021021-2202013102012310"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 003222131331 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-2300230313033210-1322013102102102-1110223210231002-0103121003122322-0133000023001123-1132022113002121-2303301213022231-1031331323200101"></a>

Type: `"object"`. single nested block, Optional.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
redirect_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303122113312202-1000213131303303-2010303330232330-2300120013032231-0222030022113211-3301302033000123-3220023000320030-1013101313202321"></a>

## Direct properties — redirect_route / 003222131331 / 3

- [headers](resources--workload--reference--group-012.md#canonical-1030302001333213-0230232331231202-0103203232233220-0031330000301230-2021212301032123-2210222030113200-2332213203131201-2131010323232132): complete subsection reference.

<a id="canonical-1300031103222230-3201322320322223-2213311310032002-1300002323023012-2102101311331330-1120030002112233-3233311132200032-1021323122120031"></a>

<a id="canonical-2233200322101133-1030002310213223-1130031103213320-1212220000303202-3201200331322031-2023001113003010-1020230230011213-2102330022001213"></a>

## http_method property — redirect_route / 003222131331 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--workload--reference--group-012.md#canonical-1030312213013310-3032021211010013-0003233111330223-3310201101002022-0123120212103011-3300101331122220-2230110110111233-2301031111011113): complete subsection reference.

- [path](resources--workload--reference--group-012.md#canonical-0110031302202213-1202221120021201-2120010113131203-0030313300331102-0331031012211122-2101232101312032-2201312123201213-0031012111303123): complete subsection reference.

- [route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221): complete subsection reference.

<a id="canonical-3032102022330131-0212313200013112-2100223323121230-0232301001320303-1032131131001122-0012112132130311-2023222123313302-2130001020311112"></a>

## Next pages — redirect_route / 003222131331 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-012.md#canonical-1030302001333213-0230232331231202-0103203232233220-0031330000301230-2021212301032123-2210222030113200-2332213203131201-2131010323232132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-012.md#canonical-1030312213013310-3032021211010013-0003233111330223-3310201101002022-0123120212103011-3300101331122220-2230110110111233-2301031111011113)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-012.md#canonical-0110031302202213-1202221120021201-2120010113131203-0030313300331102-0331031012211122-2101232101312032-2201312123201213-0031012111303123)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1030302001333213-0230232331231202-0103203232233220-0031330000301230-2021212301032123-2210222030113200-2332213203131201-2131010323232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331212333020032-0102123323330002-0133132230312123-2303031132323333-1112320101321331-0110213331213123-3013103320033303-1020010120301133"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 300210111333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-0212030122320213-1002032312133120-3200213233223001-3321210303122013-2311203011122322-2031003212210233-3020020230003020-3130320231201221"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303311211003330-1222120021132311-3120101111123113-0001031022032123-3322101031121002-2302332110100332-1130030001132020-2103011220212101"></a>

## Direct properties — headers / 300210111333 / 3

<a id="canonical-1002011033201112-0233332303231103-1313010300002013-3033132201002331-1101131020321223-0223132133212020-1332331122320202-3120031202102022"></a>

<a id="canonical-3213010130212303-3012000102211121-2211132221300003-1101130122131310-2200022013010011-2003002020223302-3032132332320330-1222310003003033"></a>

## exact property — headers / 300210111333 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1212121032011032-1033101200020013-2100212320330033-3212111310333203-3033113202001023-0000232031011331-0113130132131322-1123210333110332"></a>

<a id="canonical-1311120012032012-3332021022331311-3202113310121303-1011323223102001-1110212130333030-1310222031133231-2200010110131110-1010133130201011"></a>

## invert_match property — headers / 300210111333 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3110213323003131-3102331031210013-3300200330331012-2031302322111013-3010310010211221-3000321201231030-0132322321113230-2112010323013232"></a>

<a id="canonical-3212032303013001-0202003331310201-2112320023130320-1322301132032201-2222012013323300-0003233123021012-1111101031133233-0323202321012200"></a>

## name property — headers / 300210111333 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0031002012231203-1313023100311313-1002231322011202-3121311301220130-3133200121012210-0100310330103201-2212211130002322-3221221311020232"></a>

<a id="canonical-3123112003131321-0133311103210113-1022132212113122-1321232112102120-0303200332321212-0112121031020122-0323200313210130-1212012121121302"></a>

## presence property — headers / 300210111333 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1313211102203333-1101113301330121-3200013303331201-3210303212111031-3202010201222022-1001120212110223-3030221132032313-2210303321321120"></a>

<a id="canonical-3001011003030011-2201132111031330-1120021333112302-1112303123313331-3323000131133033-2230322221103031-3303222013231312-1112101001131001"></a>

## regular expression property — headers / 300210111333 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2132113200311012-3220321200223211-2012100103022312-0331222222303111-2233200002210102-2223333332130203-1233313302302121-1101313120012121"></a>

## Next pages — headers / 300210111333 / 9

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1030312213013310-3032021211010013-0003233111330223-3310201101002022-0123120212103011-3300101331122220-2230110110111233-2301031111011113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303312112331120-0330131320020011-1210223133132122-3131030002311120-1002031313213232-1200330312322122-3230131121300102-2012302133130213"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 011123113132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-3201122002310300-2123201333123222-0301332323111121-0333221110222213-2202022013232323-3303223310001333-2222310023322203-2211023313231102"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301331133202203-1323203301330232-2022300131000200-1213013132302300-1013113231311113-2222100100020333-3320232313310023-0222303222110210"></a>

## Direct properties — incoming_port / 011123113132 / 3

- [no_port_match](resources--workload--reference--group-012.md#canonical-3021111330110212-3033233030010122-2013220101300111-1333030320133232-2320333321321103-0131331033302030-2000032111311023-3311330210213310): complete subsection reference.

<a id="canonical-2123123221203320-1033203231323323-0300112001011323-2303002221301003-0330021010323200-2332033030131233-3213222311112012-2313322130111010"></a>

<a id="canonical-0002213330212322-0231231120132321-1201113213032022-2110330123212021-2113013102200000-2210020220220322-2002331013130112-0132301010232010"></a>

## port property — incoming_port / 011123113132 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0220232333131233-3012110203102020-1003031100321311-2210102000033130-3221232212311031-3322323001121301-3003001033203012-3021331331003203"></a>

<a id="canonical-0330131322130120-2033030223202030-0321020331102301-2020022201331001-1231302230332302-0031232210231302-1321323320222223-1220202300100231"></a>

## port_ranges property — incoming_port / 011123113132 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-3021101222320121-3302320030232203-0202211212123132-0211031221033212-0203100201310320-3120110032203332-3230202200301032-1203313231332111"></a>

## Next pages — incoming_port / 011123113132 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-012.md#canonical-3021111330110212-3033233030010122-2013220101300111-1333030320133232-2320333321321103-0131331033302030-2000032111311023-3311330210213310)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3021111330110212-3033233030010122-2013220101300111-1333030320133232-2320333321321103-0131331033302030-2000032111311023-3311330210213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013033031211333-1011211312320012-1230000120313032-1121332211302333-0200222320301033-0021130130300113-3032013203011113-3233210222032210"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 312331332203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-012.md#canonical-1030312213013310-3032021211010013-0003233111330223-3310201101002022-0123120212103011-3300101331122220-2230110110111233-2301031111011113)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0012000032102102-1230332221232030-2112320013302032-2330312230010101-0322010031231313-0331020133133321-0310020321320002-1303330233323310"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
no_port_match = {}
```

<a id="canonical-2002120313220200-2300011312333320-3112100011010322-2332111332010222-3220133112112323-3011030201321312-3110333230131221-3300321330031330"></a>

## Direct properties — no_port_match / 312331332203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303121120020131-3232233001202233-3123233021021132-2330300233321310-2103132213203222-3232330332220331-1333303122322002-0333011010303203"></a>

## Next pages — no_port_match / 312331332203 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-012.md#canonical-1030312213013310-3032021211010013-0003233111330223-3310201101002022-0123120212103011-3300101331122220-2230110110111233-2301031111011113)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0110031302202213-1202221120021201-2120010113131203-0030313300331102-0331031012211122-2101232101312032-2201312123201213-0031012111303123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200023121023111-3231010220120301-1312031110222200-3122020300310201-3232111012202131-0311303031110011-0320020320302312-3233200111101300"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 200330032313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-2011302233232110-3312011102232230-2312311121322313-3200122133031302-0023312233232121-3013110202232313-2213311232022100-3130222301331320"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322223213312311-2221311133322011-0010011103310111-1232333010320131-0212230110020320-2031301210111210-3033030012330031-3231330023313101"></a>

## Direct properties — path / 200330032313 / 3

<a id="canonical-3302303133321110-0203000201303200-1110122012220113-2203011022303200-3002101233130100-2220122203131003-3000301320231332-0123310111203210"></a>

<a id="canonical-0302132103221012-1300222332230233-2301200031020323-1011213132231010-2303123202123301-3311210311323121-2110113012212310-2003320211120223"></a>

## path property — path / 200330032313 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1032233033320301-0111213002100200-2001233210331032-1120032033200303-0003113223011301-1103121323321032-3323032010333101-1321001002323132"></a>

<a id="canonical-1020331230222111-0201021122333200-2130201323211311-0022231111031210-2123300223311000-2101312021223110-0203202320103112-0233200300333032"></a>

## prefix property — path / 200330032313 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0210033011013033-3211102232333230-1001001232332332-3002121032000302-2100301301330313-2011131120130301-0331222231221011-1232213111320220"></a>

<a id="canonical-0112323001323100-3133223023000213-2112023311231303-0003023130313202-3103001132000331-1300102112032000-2030100130202320-0220200302322211"></a>

## regular expression property — path / 200330032313 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2321210231103311-0010021010132232-2232231003233202-2122001023101131-1123002231231012-1121131130331123-1312222032202332-2220322213120233"></a>

## Next pages — path / 200330032313 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011023121203020-3332123201012212-0010000232322201-0131212302221121-0310320100120012-2122000230103203-0301103120121313-0111323031030213"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 000102211323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-0203331000300010-3101311102012313-2131333223212021-3232212023120301-1111122023001121-1112230221210011-1023120322310323-1011013211021210"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010002212211313-3213300332211330-1302320000113232-3022321102030001-0203121332102212-1211302112211032-3321200003231100-0310301021103200"></a>

## Direct properties — route_redirect / 000102211323 / 3

<a id="canonical-0331313101020130-2013030021002120-2102021332222113-0320001302301132-1313000033032300-3123220111103021-3232122300231233-3221120301001002"></a>

<a id="canonical-3032013212323202-2013033332213322-1032222020320132-0033333312301210-1200230212321202-3223233313021122-3202220232132033-1113033203210102"></a>

## host_redirect property — route_redirect / 000102211323 / 4

Type: `"string"`. Optional.

Swap host part of incoming URL in redirect URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033202111030102-2021321031103230-3002200113311232-1330101233222032-0000330111011110-3320311221331212-3003120032233212-3222221131211303"></a>

<a id="canonical-2103201210132020-1031111010310120-3330210213211321-0121330323002211-3101100322301013-0131101312112322-2032102112101220-0030212300101011"></a>

## path_redirect property — route_redirect / 000102211323 / 5

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0310321121122213-1022221123132022-0031031113001233-0101101230302000-0103101220201031-3023303321123232-0313131212120020-0020210002030000"></a>

<a id="canonical-3313203122301222-0031331111122230-0330303031000003-2311322033021101-1230131012013312-1223311031213320-3020010331131223-3221122232012003"></a>

## prefix_rewrite property — route_redirect / 000102211323 / 6

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2330103220013133-3331220322101222-1000221003300311-2211220122303210-2313102121101021-3233003211120211-2120233021121132-1002010023113113"></a>

<a id="canonical-0333212321012122-0033321120011132-0201003020101220-3323122133003122-1223022312331220-2022320102301011-0321222330103303-1111300001100131"></a>

## proto_redirect property — route_redirect / 000102211323 / 7

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](resources--workload--reference--group-012.md#canonical-3213010130011233-2222101033320310-3112210212110200-2231011121301231-1030030302030000-0333301200123312-1322000131313301-0131131223201011): complete subsection reference.

<a id="canonical-2131302011332321-2232333011213210-3322133320121010-0200122021132231-0201102001313313-1332011321131231-1322211331020310-0003000310011320"></a>

<a id="canonical-2311323012300312-3330302031001221-2111110023030033-1031202220211122-0301310111322130-1203312030332221-3103033111203201-3022102031011320"></a>

## replace_params property — route_redirect / 000102211323 / 8

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3012031331103310-0012120322321102-3010220210210333-3122312103201103-1010012001000011-0002300031130211-2103022211231003-1032102011121213"></a>

<a id="canonical-2031120310011022-0131021130310101-2130033130023102-2223103113210330-0313133120110110-3310131131230023-1013103230000013-1231220330023322"></a>

## response_code property — route_redirect / 000102211323 / 9

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](resources--workload--reference--group-012.md#canonical-2202221221020330-2011232312031133-0133231220113200-2023020103323123-3012333331132001-2131003223011211-3331203312231112-2203100210022112): complete subsection reference.

<a id="canonical-0231312130302103-3010130133232200-3113003132131210-2010231131133213-3130122100103203-3101032201333220-3221122211333232-0202213233102232"></a>

## Next pages — route_redirect / 000102211323 / 10

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-012.md#canonical-3213010130011233-2222101033320310-3112210212110200-2231011121301231-1030030302030000-0333301200123312-1322000131313301-0131131223201011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-012.md#canonical-2202221221020330-2011232312031133-0133231220113200-2023020103323123-3012333331132001-2131003223011211-3331203312231112-2203100210022112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3213010130011233-2222101033320310-3112210212110200-2231011121301231-1030030302030000-0333301200123312-1322000131313301-0131131223201011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110100022332023-3303033322102303-2103220233020333-3100301230300121-3110322303112303-2022313001200010-2122331201112311-3131100110220232"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 121313333311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-0133102132211203-3313003320010231-3313120110022312-3033232221021203-1322301332311200-0103213020222132-2232203223200313-1121030123232203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
remove_all_params = {}
```

<a id="canonical-0330232301122021-0221012301032220-1102132200301200-0020212100221331-3002002031232012-0033002021201231-2012133111030111-3113330222110213"></a>

## Direct properties — remove_all_params / 121313333311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332332300133002-1033312312221232-1321200233233303-2232310001122221-3011102012333321-1131123332330300-3032323013020300-1010022030312231"></a>

## Next pages — remove_all_params / 121313333311 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2202221221020330-2011232312031133-0133231220113200-2023020103323123-3012333331132001-2131003223011211-3331203312231112-2203100210022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111220100020023-3231301031000003-2002133310211203-0231330331203230-1000232323032222-1220221101121230-0222101111330220-1102011220011232"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 313200010032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-3232202130133333-3233010022102012-2101020122102200-3310012002330203-3121122110232013-2023120031022131-3122230333230121-2303221233113002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-1023122111223323-0030210032233111-1132232301101130-2211322103122203-3320100122010211-0201231020122000-3320103320320011-2120111320323231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
retain_all_params = {}
```

<a id="canonical-2330010300123100-3130102221001003-0330122032111000-0022230102122001-1002201102230313-1110001310121230-1213013332012230-1111231103233330"></a>

## Direct properties — retain_all_params / 313200010032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132021330013200-0120031023012313-0101203210023033-3122100132111321-0132210322132103-0313313210000133-1130120333221112-0111233210121302"></a>

## Next pages — retain_all_params / 313200010032 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-012.md#canonical-2001031101031300-0320210021331321-3321323300033002-3223100301221103-1033230103201322-3100320020322030-1032113103032023-3303220132000221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103133202303133-1013201231222233-3210002301220031-0030300230032100-3223233212100201-2300233013002012-0313212122121202-2333222112231320"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 222032101113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-2100130010110001-1230303123103131-3332322310021233-3121220212000211-3012310032002001-2011123030001003-3222221101312003-1303301231022112"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010221012012112-0302312220322322-3330132123002301-3103200020220211-2211021020103102-0022313333301021-2010203320133202-2321031220202103"></a>

## Direct properties — simple_route / 222032101113 / 3

- [auto_host_rewrite](resources--workload--reference--group-012.md#canonical-0032210102023012-2023123122022322-2313131221032111-3302222313320332-0333322331112201-0321101112221133-2002310131112222-2321021302220121): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-012.md#canonical-1100222233110030-0231100312330000-3303310013120330-0013323300103321-2103023123321122-3330233323023302-1210102230123002-3001222333032002): complete subsection reference.

<a id="canonical-0112302331222002-2121303112210120-0323300100331202-1012212012020110-2221333221012303-0131133300222323-1030232222032003-2011300131331002"></a>

<a id="canonical-3011122202130033-0120233132010110-2120103111202230-0031231102012310-2101213323033021-3231102101123121-2011331130212211-3302310201222300"></a>

## host_rewrite property — simple_route / 222032101113 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2310310130322000-1101003003330100-3311230233132112-2210232120013213-0310102032221203-1000312223321132-3331323221123123-0233301230300031"></a>

<a id="canonical-2210301013323111-1132132120322222-0323021210100133-3103021312202230-0313002130323211-3132102212031203-1210103113011232-1332230302232113"></a>

## http_method property — simple_route / 222032101113 / 5

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [path](resources--workload--reference--group-012.md#canonical-2331223302320323-0223101203023220-0022233233203133-0011031313020213-0122132320321011-0102122220033001-2232302013032131-2113331232313331): complete subsection reference.

<a id="canonical-2220213023032321-2000310123031210-1012323022310101-2313230120330103-3130103232203003-2113213020033000-2030121113213303-1233123301312113"></a>

## Next pages — simple_route / 222032101113 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-012.md#canonical-0032210102023012-2023123122022322-2313131221032111-3302222313320332-0333322331112201-0321101112221133-2002310131112222-2321021302220121)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-012.md#canonical-1100222233110030-0231100312330000-3303310013120330-0013323300103321-2103023123321122-3330233323023302-1210102230123002-3001222333032002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-012.md#canonical-2331223302320323-0223101203023220-0022233233203133-0011031313020213-0122132320321011-0102122220033001-2232302013032131-2113331232313331)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0032210102023012-2023123122022322-2313131221032111-3302222313320332-0333322331112201-0321101112221133-2002310131112222-2321021302220121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311023301133321-2002311001211011-0113213123031122-3132021112123333-2012303001031301-1132112210302333-1003132000101332-1320202322111303"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 301111302320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-2033102221132301-0133322332310310-2313023210221122-0122230133122213-2212230111130102-1112013333103222-0120301210231213-1011330201221233"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
auto_host_rewrite = {}
```

<a id="canonical-3102122300002202-0133012022322213-1021023010123131-2032013300203300-2233220133113312-1100113122301033-1313220303033303-3320300233013131"></a>

## Direct properties — auto_host_rewrite / 301111302320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103221101103230-2010313113310131-1200032311021000-2000220111331020-2202122333022012-0000021011123210-0323022131133310-2230331132000100"></a>

## Next pages — auto_host_rewrite / 301111302320 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1100222233110030-0231100312330000-3303310013120330-0013323300103321-2103023123321122-3330233323023302-1210102230123002-3001222333032002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300331333102313-2312111212101300-3231023111102132-1330110011120220-3002020333230203-2233221323002233-3213200001100120-0202310110030220"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 002201230100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-1302231312030012-0030210233331132-0000203320130101-2030032310112133-2333223021211113-1120131213033332-0311301223011203-3031230101023030"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_host_rewrite = {}
```

<a id="canonical-2230211313020010-1332332111322013-3221303013132133-3301212101010003-1120110300320100-1003000002313320-1011213002311022-1031213102122220"></a>

## Direct properties — disable_host_rewrite / 002201230100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003212211310100-1312023201112320-1133321201120122-0200010333301312-2222002021012131-1022200221331222-1000310323223012-1302031233300103"></a>

## Next pages — disable_host_rewrite / 002201230100 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2331223302320323-0223101203023220-0022233233203133-0011031313020213-0122132320321011-0102122220033001-2232302013032131-2113331232313331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223222223001202-0112312222002323-2130210203022012-0333121231002302-1030112012330000-2000102010022123-1031122332002331-0031001102233303"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path — path / 121031203123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-2233101221201201-0221103033201012-1312003201000130-3330220121223312-0322131110123233-1130032013333000-0312333002333002-1322131130200022)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-3111001331002333-2213011130302122-2200001110221002-0331213233222233-0233122202123131-0213121221212021-1320223000322131-2121211010211202"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031001321333313-0102123133311132-3321310313222332-0330123002113221-1221231300213010-2001130200121310-0013032230231233-2232312001320211"></a>

## Direct properties — path / 121031203123 / 3

<a id="canonical-3111302232221023-3013000120013211-3101323212332001-2101221313211310-1333111031210312-0321301132112221-2122313321103000-3333203222211020"></a>

<a id="canonical-1001102131303221-2333003012223323-0302120311100222-1100323303312103-1220023212023313-2321323121012123-0113320230120333-1013303032301301"></a>

## path property — path / 121031203123 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3320213231311210-3320032000201232-1020000301200020-3112101101020332-0101002022013202-3010022220033112-0111333111310031-1231312221112212"></a>

<a id="canonical-3001311301133332-0212010111132012-1132213213133030-3211112230312122-1120131220033112-1001311211312001-0212310323311332-1002332303130033"></a>

## prefix property — path / 121031203123 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3131103102112002-2013032030320313-2012030022102033-1303301112220123-1000321222213102-2112030333010003-2311030301312003-2200221230013320"></a>

<a id="canonical-1000030210303321-1020202132121010-1332232213223010-1103210201200330-2103330322323313-3323001013322302-3213302312132312-0010002101103200"></a>

## regular expression property — path / 121031203123 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0311120201231233-3203130123012001-0023002123213213-3120100130033221-2331011001110013-1301102333130130-2003213101000323-0332321002000301"></a>

## Next pages — path / 121031203123 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-1012203321312032-1210001322011003-1210311231233020-3110300033300022-1313221111233030-3333033300210133-1230313311211221-1223332020223123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112003230010132-1021130213020201-1121111010320322-1211110023132002-1110013301030012-2301110300121301-0123020233332122-1303100210102301"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port — port / 132310223033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-3001030333111132-0213023220321321-0132310212031012-1201313012122331-0311200112211203-1012022031301030-1301023233300130-1300010002230220"></a>

Type: `"object"`. single nested block, Optional.

Port. Port of the workload.

Upstream description:

Port of the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210323031112110-1030301033123230-1332131000213113-2021303203330000-1001233000113101-1301030201230210-3003230101233103-2010003003331201"></a>

## Direct properties — port / 132310223033 / 3

- [info](resources--workload--reference--group-012.md#canonical-0022230033210322-1030012030011331-2011320103030321-2313030221000123-0102012221332322-1122111212231213-1013031321030313-0021002102112323): complete subsection reference.

<a id="canonical-2000312313203333-3131133100011322-0220301212213032-1301323030031210-0030111210131123-3011023003312030-0010303031121222-3300101230221023"></a>

<a id="canonical-2221010321103211-0112001311110210-2321300223313312-3333011010100203-1321300221102333-3120112003323231-0000201011133323-2113321322230021"></a>

## name property — port / 132310223033 / 4

Type: `"string"`. Optional.

Name. Name of the Port.

Upstream description:

Name of the Port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1033013230320000-1311202000303302-2010101302111000-1301323111213313-1131233312213321-1030113031103120-2210012000301010-2331323010211112"></a>

## Next pages — port / 132310223033 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0022230033210322-1030012030011331-2011320103030321-2313030221000123-0102012221332322-1122111212231213-1013031321030313-0021002102112323)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0022230033210322-1030012030011331-2011320103030321-2313030221000123-0102012221332322-1122111212231213-1013031321030313-0021002102112323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303222101222010-3101211120312332-3100020313012100-2122023221330203-0020032000001002-1212232112200330-2030003202231231-0312103033331202"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info — info / 021100113130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-0222303313123200-3032021132112000-1010201221303323-0133132023232331-1312102013121302-1301213301132112-2010213023231303-2032230201101303"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200332213230033-0012112031222333-2213132232113100-1120233010300021-3322323001032210-0012123300010213-2012213221020232-0200021103103102"></a>

## Direct properties — info / 021100113130 / 3

<a id="canonical-2020310020001112-3221030222032202-0303002023301230-2113013313221031-1101232031202112-2013201021110132-3133303030312322-1230300301222100"></a>

<a id="canonical-3033310321022232-1331133131013113-3232311110210102-2001111223111101-2110312030203103-0212321302032110-0202001313220203-2211133131102121"></a>

## port property — info / 021100113130 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0312132020102310-1033310023202133-2021332100112230-0312111322112302-2322202121222020-2221010132121231-2001021003332213-1220301131231210"></a>

<a id="canonical-0320122210212100-3030210113002311-0110200012121310-1333001113203133-0333301121301211-2102312200013030-0010012000011012-1323321110032012"></a>

## protocol property — info / 021100113130 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](resources--workload--reference--group-012.md#canonical-0302301220132301-1312111303001102-3003100030302011-2123010000300002-0120323133321123-3221112012320110-3302231210110023-0123231212130011): complete subsection reference.

<a id="canonical-2001203202203120-0213213201110123-3011111010302310-2200311322000302-1011331321100013-1232202220112123-3131000103033213-3133333312130332"></a>

<a id="canonical-2121333111303120-2031213330212031-1331232033333211-2120120231232002-0122033110223312-3210322000000013-3230202311010302-2111030203222122"></a>

## target_port property — info / 021100113130 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0200123132001022-3213002302200213-2233212222310000-2322202301101130-1321112120130211-1013322112212330-2330030002211002-2013103222112212"></a>

## Next pages — info / 021100113130 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](resources--workload--reference--group-012.md#canonical-0302301220132301-1312111303001102-3003100030302011-2123010000300002-0120323133321123-3221112012320110-3302231210110023-0123231212130011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0302301220132301-1312111303001102-3003100030302011-2123010000300002-0120323133321123-3221112012320110-3302231210110023-0123231212130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003332200012033-0113201313201333-2112330210100122-3131002202302312-1111330011103223-3033223110023332-0232233110132130-1031300201330221"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — same_as_port / 233102300230 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0022230033210322-1030012030011331-2011320103030321-2313030221000123-0102012221332322-1122111212231213-1013031321030313-0021002102112323)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-1122322112132223-0133233121312313-2122312112232132-3131033022100123-3130012022122303-3331110310002103-2210130220123310-1103010031232330"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
same_as_port = {}
```

<a id="canonical-2111232000321001-0131202101012322-2012211013023200-3211323233103231-0232211322030010-0103011013000233-3102102231110110-1220220233010113"></a>

## Direct properties — same_as_port / 233102300230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332102211130220-0202010030212313-0023021331201331-1330331232210212-3321011201212233-2232233013322232-2201222010112230-1022112031313032"></a>

## Next pages — same_as_port / 233102300230 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0022230033210322-1030012030011331-2011320103030321-2313030221000123-0102012221332322-1122111212231213-1013031321030313-0021002102112323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1201000230031032-2312033000023331-1111332222123100-2001203111302231-0133102032000130-2121122320101311-2223311020113123-1220222002313220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213013212300013-1311010102323122-2003130232230230-3310112000012333-0313202203213112-1313220220030200-1002131012011123-1003321210101131"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — tcp_loadbalancer / 331120301033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-3301133231110221-2002121212013113-2301230102320201-1231233302132300-2201101203033210-1022300112202100-2222001102111020-3310330233222122"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020223120111110-1223302330230201-2231111003232323-2122332112030101-1233232323330321-2001323211210022-0312123010102013-1031022110331120"></a>

## Direct properties — tcp_loadbalancer / 331120301033 / 3

<a id="canonical-3223103120332110-2100132130021000-3010131210131200-1221001120120220-2233132200220113-0303120202300213-0300120102130120-3220233031331303"></a>

<a id="canonical-3333320321103231-0301231133203312-0323200030203000-1212112230133102-2210331000313113-1210012222233011-2313320101033302-2201312213123013"></a>

## domains property — tcp_loadbalancer / 331120301033 / 4

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302101021103103-3321201101323231-3212120301303221-1212213110200033-0331132010201131-1113332333333213-1303201233030002-2210221121311011"></a>

<a id="canonical-1002300020220000-2303302033310132-2310031201122220-0330231021221320-2101202320111312-1321111111302121-1031202302123303-0230111311121101"></a>

## with_sni property — tcp_loadbalancer / 331120301033 / 5

Type: `"bool"`. Optional.

Set to true to enable TCP loadbalancer with SNI.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233123013121313-0312120321031321-0301230313123320-2323302000310120-3122321321130002-2202123103132331-2202310232210103-0000200012101212"></a>

## Next pages — tcp_loadbalancer / 331120301033 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102332210121133-2123001322113021-2213222211020023-1301311031322023-1000222212320132-0310003231112211-1011130132311111-3301213332332322"></a>

## service.advertise_options.advertise_on_public.port — port / 200322111331 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- service.advertise_options.advertise_on_public.port

<a id="canonical-1120220132203311-1233230110131322-1012113213001013-0301101200022122-3323302122212220-3112020010311121-1320031221321330-2132023123330302"></a>

Type: `"object"`. single nested block, Optional.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201232010231113-1322013110300231-0001231110311010-3021011320002210-2021032132002100-2122331212220332-3301021212002123-2323321220110302"></a>

## Direct properties — port / 200322111331 / 3

- [http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330): complete subsection reference.

- [port](resources--workload--reference--group-015.md#canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-015.md#canonical-1002130311113332-2130231123003000-2113012220322103-2102212131300021-3233123213121332-3000203213103023-2322231100121132-3001110212313022): complete subsection reference.

<a id="canonical-0233112023333112-2131220023013330-1111121133210320-1212213013223232-2021011101130303-1103212201221311-1331232323003320-3101221211130103"></a>

## Next pages — port / 200322111331 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101)
- [service.advertise_options.advertise_on_public.port.tcp_loadbalancer](resources--workload--reference--group-015.md#canonical-1002130311113332-2130231123003000-2113012220322103-2102212131300021-3233123213121332-3000203213103023-2322231100121132-3001110212313022)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013033103111211-0233010203302302-3101320211130112-3320133333010300-3232103223102230-3312330011010030-2203222223332132-2232220312210012"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer — http_loadbalancer / 302023303231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-1222020020311031-2230103000211330-1013100322022031-0302333303333330-2130003000323320-3231333213110333-0222133123330033-1101231011131100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223010301020012-1332222101020212-1332032133130231-0321130332300103-1233112010212013-1313311302030113-1103013133031213-0312103120222202"></a>

## Direct properties — http_loadbalancer / 302023303231 / 3

- [default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013): complete subsection reference.

<a id="canonical-0000120110103220-0100113133310230-1001213123320003-0211232121003020-0132303232032120-2330110212211333-1112231021000322-2033011321202230"></a>

<a id="canonical-3323201011201000-2212332321013120-1030111122221022-3020021101102123-1302313000213032-0012212223301033-0220322030203310-3100203011033232"></a>

## domains property — http_loadbalancer / 302023303231 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](resources--workload--reference--group-012.md#canonical-1320030113012233-1022032312211012-1001221123023230-2222203332110111-0020322321131123-0203312222011111-2231002231202012-3102013312011000): complete subsection reference.

- [https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033): complete subsection reference.

- [specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311): complete subsection reference.

<a id="canonical-0102230312102212-0001113023133201-2130012020322102-0132132311310213-0221212330013132-2010233322322310-1001032230113112-0300222033131310"></a>

## Next pages — http_loadbalancer / 302023303231 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.http](resources--workload--reference--group-012.md#canonical-1320030113012233-1022032312211012-1001221123023230-2222203332110111-0020322321131123-0203312222011111-2231002231202012-3102013312011000)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021122231332222-2333213103201033-2003203120003221-1131120222032000-1133120232023301-2202030022202233-1223121032121303-0323301003301131"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — default_route / 013211320101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-0113023020203113-0133202311230312-3213022303012213-3100201110002221-0100302233221102-2111221232020301-3100002020210331-3230211021130121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232010030022322-2301120213030101-3113120320110121-1133012013333323-1002220023131232-0223032112330303-2312033332113010-1001203020322101"></a>

## Direct properties — default_route / 013211320101 / 3

- [auto_host_rewrite](resources--workload--reference--group-012.md#canonical-1033212013303333-2332332121233032-2233233303102211-0132302110133202-0231110022013023-3112213012303231-1323202023200023-1102301131011112): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-012.md#canonical-3110022032312000-2322311000113113-1320103313311301-0120131111230013-1013123000221212-1030330020111312-2221131111222002-0201210321013213): complete subsection reference.

<a id="canonical-3012100013221011-2023331201311230-2100221111130123-3101313012003010-0033300202120130-2130202213131112-0210011130130330-2011030203200030"></a>

<a id="canonical-0000022222321032-0323030233103131-3113111331020233-2301133012310333-0020202023022300-0332203022132201-0303031220213001-0312220300000020"></a>

## host_rewrite property — default_route / 013211320101 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2101031122303113-1103320232331023-1313213211200330-2122110202110031-0122130002223201-1321100311310301-1112121211000211-2212303000120103"></a>

## Next pages — default_route / 013211320101 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-012.md#canonical-1033212013303333-2332332121233032-2233233303102211-0132302110133202-0231110022013023-3112213012303231-1323202023200023-1102301131011112)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-012.md#canonical-3110022032312000-2322311000113113-1320103313311301-0120131111230013-1013123000221212-1030330020111312-2221131111222002-0201210321013213)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1033212013303333-2332332121233032-2233233303102211-0132302110133202-0231110022013023-3112213012303231-1323202023200023-1102301131011112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111210312010113-1301323113121322-2010202030332222-3101133201133231-1203220103122212-2230201030221002-0321133311222332-0120300131200132"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 211232101022 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2213201333111223-1302220003223030-0122322013101030-0311123012000102-3212330102322130-0201021303113210-2002101222001220-2330102233021120"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
auto_host_rewrite = {}
```

<a id="canonical-2233033220212103-0030221111222122-0210011212210331-1103211100332012-2233112301123330-0132100200023112-2202113010122113-3100303032230122"></a>

## Direct properties — auto_host_rewrite / 211232101022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331102023300002-3013212030200123-0320221213233021-2103323210302012-1310131022023131-0030331101100331-1021300032123230-1112203000101110"></a>

## Next pages — auto_host_rewrite / 211232101022 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3110022032312000-2322311000113113-1320103313311301-0120131111230013-1013123000221212-1030330020111312-2221131111222002-0201210321013213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230001211002323-0211303131231313-3002330131312120-0030210000212112-3222003101130313-0312221202320122-3301003030100231-3200223002130002"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 202213303011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-3302012021102300-2011330032302323-2100001333223303-1313301112100132-3121323232121020-1312110110223031-2103123310102302-0200020202332200"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_host_rewrite = {}
```

<a id="canonical-3220011332123003-3312100112220031-3220313231131300-0001001010032311-1123203212023200-3310230233030113-0333321331103112-3131302121112202"></a>

## Direct properties — disable_host_rewrite / 202213303011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112130320230130-2031021020320020-2202033303113110-0010010130030033-2201211120123123-3211313022332121-3010022120132303-1132020322003133"></a>

## Next pages — disable_host_rewrite / 202213303011 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0022310121100111-2332031122200103-1033031102121220-0103333213311021-0111300321332202-2003101303011032-2031332203232303-0333120212103013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320030113012233-1022032312211012-1001221123023230-2222203332110111-0020322321131123-0203312222011111-2231002231202012-3102013312011000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102210300110113-3110031301220222-3122213013333023-3223021221120030-1323332130111300-0332202021001113-1333110110113130-0210310313000233"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.http — http / 022112230011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-0031323223002202-2331221023220310-3230101122010002-2013231323222002-1021232213230322-2200232013201010-3133131323122233-0032231203301330"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302122232131130-3202132220012031-1002330122030102-0312211002112032-1101232323310232-3323330103331003-1022301120021200-3322211012322301"></a>

## Direct properties — http / 022112230011 / 3

<a id="canonical-2022110331032023-3322203310312202-2011220323021331-1020201201222113-1121032320000103-3321232331313323-2210133233203231-2021120220203323"></a>

<a id="canonical-2332200110133030-0232132332203121-0032213221002032-1112023033003221-0102000211321030-2303030120003003-0121120001313233-1001010021232331"></a>

## dns_volterra_managed property — http / 022112230011 / 4

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200023213213032-0313332222222111-1301203223203330-2232103011302323-0100333330233113-3001021100200031-2203233232232113-1113220233220023"></a>

<a id="canonical-1003220313333002-2133132123322122-1200103230210232-0032222313123212-3301122211102011-2233120113012010-1311100200320221-3332130313320232"></a>

## port property — http / 022112230011 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0323300120311121-1303132220103203-1030231320231002-3203300313323002-0233321313201111-1130010102200120-0003211011003233-0012011203120213"></a>

<a id="canonical-1232302021101331-0223011321103300-2212231031223331-1220032020303032-2030303203022230-3320312232322210-1333100000132011-1131130233231302"></a>

## port_ranges property — http / 022112230011 / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2202303033310302-0210120203111300-2232111213221031-0200113010111323-1011021122301200-1323003222101012-0023303001101000-0301110211320213"></a>

## Next pages — http / 022112230011 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300312333100200-1320332133000212-3311313023132012-1312212000000332-1033213322002102-0022130121310302-3230022303101121-3132210212113123"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https — https / 212333301333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-2323030002233233-0202303330123212-1312030120031331-2212130310220101-1310013333033032-3213333002010230-0330122021123221-1323113333302131"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220123000033033-1320233311322003-0111202132331223-2312122123322100-3021033132130031-3132013022320132-1212230022030323-1303220301223222"></a>

## Direct properties — https / 212333301333 / 3

<a id="canonical-0212012023333000-1310321110323210-1100103330131330-3322103010030012-2223211311312121-2020202133130022-1012200222322122-2103312233032202"></a>

<a id="canonical-1211022330113313-1322013002300202-3302121123220320-1011222002101131-0102221130110313-1112303322022031-3211020320331203-1130221100212111"></a>

## add_hsts property — https / 212333301333 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0103233010332103-3130122311312311-0211002100312231-0113122232233110-0012103312322331-0131001121300130-2112230130032232-1210131211110333"></a>

<a id="canonical-3210310113120233-3023102223330312-2013112000322230-2020200233211320-0230010200303112-2230102133012023-3121220123221300-3302233032220100"></a>

## append_server_name property — https / 212333301333 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220): complete subsection reference.

<a id="canonical-2331333301130023-0202102221102332-1112112130200011-2233113133230213-1220133210120021-3331233030022013-1002211022232210-2201100120301313"></a>

<a id="canonical-2203301100203103-2100111331020122-1223013321331003-0223232110312122-2310120210030110-3023303133212231-2211032300221121-3033231330103313"></a>

## connection_idle_timeout property — https / 212333301333 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-012.md#canonical-3321310121130013-0222122313201310-2213231322223211-3200100220211010-1310022022201332-0101103120113210-0221322120201202-2300120220310301): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-012.md#canonical-1320132313310100-3321223020113201-0130032123021103-1131232332230310-3230213202011331-3201311113003121-3211003013011113-1033213131023001): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-012.md#canonical-2221311312002300-1300300121022222-0232321211132023-0311332020111332-0022132113301000-3032221303002122-1123313231221231-1200113120011331): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-012.md#canonical-3233203303111321-3231322112033101-2221200310110223-1202313122212102-1223331231202221-2121122131101331-1010313202012022-2203102032221023): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-012.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123): complete subsection reference.

<a id="canonical-3001021302323111-1121110212031203-2333221133332010-0123122232333320-0312103321323113-3323320012200332-3212000123333021-1323300222132132"></a>

<a id="canonical-1003010123131000-3232302321123130-1130002102011101-3211020121332320-1321113032221032-0111032203220332-3220332301313020-3102311110332302"></a>

## http_redirect property — https / 212333301333 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](resources--workload--reference--group-013.md#canonical-0201012102332122-3021012133330000-1212133112031131-3000111023333013-1113210313110021-3303133302312331-3303212013301222-0032301130003232): complete subsection reference.

- [pass_through](resources--workload--reference--group-013.md#canonical-1331231121223122-3020333331231320-2222002002323013-0001220313222302-3223121012120332-1013121312110111-0010212110113013-0200301020223132): complete subsection reference.

<a id="canonical-1221312012021113-2113012111120323-0111122320323111-3322130301332232-2013001222033133-1100110230233313-0021003312103233-1231322111233210"></a>

<a id="canonical-0103202321303323-3221212100233033-0001302231223322-1203031003101021-0231032123031201-2001232333011032-0130230030001211-1133023031132211"></a>

## port property — https / 212333301333 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0020010320031323-2320102133132133-3011112100031030-1331213323000212-0020023101211221-3331213010111021-3300000212221120-0021300031121300"></a>

<a id="canonical-1312112022103220-2202220110003203-2101332321201033-1333312013110111-2113131121000013-1103130300330000-2133123300000210-2323033011322031"></a>

## port_ranges property — https / 212333301333 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1210323300033212-2202121030312121-3220130311313001-1222100013233030-0303112330030301-3330130203223223-3332121011013002-3230311102223213"></a>

<a id="canonical-0022130301110121-3211011321202311-0213130111322323-3031013231101330-2332233120320203-0121102122212030-0323212023220230-3213100130132331"></a>

## server_name property — https / 212333301333 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310): complete subsection reference.

<a id="canonical-3322220202303120-2102322010200031-0011012030210312-3210121321122122-1233201333210111-0032010230112222-2301203213320203-2113310331213002"></a>

## Next pages — https / 212333301333 / 11

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](resources--workload--reference--group-012.md#canonical-3321310121130013-0222122313201310-2213231322223211-3200100220211010-1310022022201332-0101103120113210-0221322120201202-2300120220310301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-012.md#canonical-1320132313310100-3321223020113201-0130032123021103-1131232332230310-3230213202011331-3201311113003121-3211003013011113-1033213131023001)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-012.md#canonical-2221311312002300-1300300121022222-0232321211132023-0311332020111332-0022132113301000-3032221303002122-1123313231221231-1200113120011331)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-012.md#canonical-3233203303111321-3231322112033101-2221200310110223-1202313122212102-1223331231202221-2121122131101331-1010313202012022-2203102032221023)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-013.md#canonical-0201012102332122-3021012133330000-1212133112031131-3000111023333013-1113210313110021-3303133302312331-3303212013301222-0032301130003232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](resources--workload--reference--group-013.md#canonical-1331231121223122-3020333331231320-2222002002323013-0001220313222302-3223121012120332-1013121312110111-0010212110113013-0200301020223132)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231022031210221-2132223013011320-1332100023002100-3132322231330332-1020201013131302-1030113321210000-3330333333210213-0202102102033101"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — coalescing_options / 123301321300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-2122121130121203-0321200021221233-0223331132131030-1020322011310321-3033000311102311-3313110211130202-0002320202230111-0111212110301122"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030222303330033-1031122330113111-0002202200032330-0223031132211200-0210231033323310-2131021221211033-2320322303102220-0132333331131300"></a>

## Direct properties — coalescing_options / 123301321300 / 3

- [default_coalescing](resources--workload--reference--group-012.md#canonical-3102301230021131-0233123330123231-3033332321113121-2100323030123023-1101030031021131-3311221233310211-0332320212031033-3032100121121121): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-012.md#canonical-2212331311020211-3011111311300012-3012210221213023-0103033200211213-2102032221201321-0123023212112023-3323000311303320-0100230011033011): complete subsection reference.

<a id="canonical-0102133221220203-2120103323103103-1222013333110223-3112212213303310-2010201030012011-1020200030001223-0021303013013200-1323300320110030"></a>

## Next pages — coalescing_options / 123301321300 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-012.md#canonical-3102301230021131-0233123330123231-3033332321113121-2100323030123023-1101030031021131-3311221233310211-0332320212031033-3032100121121121)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-012.md#canonical-2212331311020211-3011111311300012-3012210221213023-0103033200211213-2102032221201321-0123023212112023-3323000311303320-0100230011033011)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3102301230021131-0233123330123231-3033332321113121-2100323030123023-1101030031021131-3311221233310211-0332320212031033-3032100121121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031013210113311-1001001132313000-1223221032123133-2221001323023033-1010300201110213-1221110103322123-0030311101331022-3112012321012010"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 221022230101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-2101300123311302-2121001131332011-2323220003301222-2301120332003201-0233122230321310-0122300233333000-3132311301002231-2213130121033112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-1312203011212223-2333211230121211-3000201101331213-0312120323111331-3003330310010033-3213320332200133-3222011301122311-1121123133211312"></a>

## Direct properties — default_coalescing / 221022230101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213320311231310-2001003223023300-1101331002311011-1032222313010223-1031010222131032-2213031301132303-0113212301320112-0020032121131303"></a>

## Next pages — default_coalescing / 221022230101 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2212331311020211-3011111311300012-3012210221213023-0103033200211213-2102032221201321-0123023212112023-3323000311303320-0100230011033011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332201201311200-3302123331101001-3112312133223001-3133131332122220-3013113002103001-2202320210213200-2122020212210103-1032313000033001"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 132033032123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2030113112032200-2111122320102223-0213232300131302-1210211133203233-3132322122120133-0211330202323201-1032202132203212-2212022100022010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-0012112111010021-2113232110231312-3313132123033221-1301221300020032-3110322023333201-2212302022230210-3021123311122323-2203033101232110"></a>

## Direct properties — strict_coalescing / 132033032123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123121003123322-1211330001110321-1012123322220331-2031020000322131-1223021220003133-1010103112131130-1002003003101332-3130321333130202"></a>

## Next pages — strict_coalescing / 132033032123 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3321310121130013-0222122313201310-2213231322223211-3200100220211010-1310022022201332-0101103120113210-0221322120201202-2300120220310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330300302330312-3231133000113121-1213222123020230-2333123112312112-1020200310000331-0232030031210103-2123020230112233-3312302012332211"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — default_header / 111212002312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-0013231333303232-2031100303331320-0302231322300032-0331100111103323-2101032013101213-2210022103023323-1012003332000313-1113132123331010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-0120330130003322-0233101223030023-1331000312211230-1032121301121012-2320111003221222-2031112012202003-3331131231000123-0131222001311112"></a>

## Direct properties — default_header / 111212002312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311103212301322-0113323210123303-1320321001222122-0032101011122201-1323112302213232-3201133220032112-3003332122312301-2322311102003023"></a>

## Next pages — default_header / 111212002312 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320132313310100-3321223020113201-0130032123021103-1131232332230310-3230213202011331-3201311113003121-3211003013011113-1033213131023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310310113002001-1110201011020233-3203302330323002-3112113222211123-0223312212021002-3202233003032301-1100112332002213-2221030011200302"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 111300322211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-0001202120301200-1032111230300210-1033222203030230-2331313312032003-3310100323232320-2031331131022132-2002030330302302-1301112233003230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-1120203333010131-1232332300102212-3232021202201102-1103230222020120-3231332230012331-0103203320333301-0030322012220211-0321032133100131"></a>

## Direct properties — default_loadbalancer / 111300322211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130110232303330-2033012102030013-2220033223022211-3120331221103003-3120101220201220-2003010322131233-1302021230202330-0030000323021000"></a>

## Next pages — default_loadbalancer / 111300322211 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2221311312002300-1300300121022222-0232321211132023-0311332020111332-0022132113301000-3032221303002122-1123313231221231-1200113120011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121033031013303-2123133201122223-1112103233211321-2233200010101333-1003131210112111-2210300223101321-0013202203330130-1131320013232322"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 131330230330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0301122012020302-0332113222320110-3301220322221002-3130120310010222-1020321103320202-3222120112233202-2011222123230111-1001200311121121"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_path_normalize = {}
```

<a id="canonical-2220212121210202-0222123330030010-0330111022310030-1311033213301232-0302032233201333-3103121200211202-0030233323322301-0121300010322223"></a>

## Direct properties — disable_path_normalize / 131330230330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113012203021020-0031122113112213-2201102213321022-3002310130221303-3221323123112110-3012211010221211-0121010003020131-2233232200211312"></a>

## Next pages — disable_path_normalize / 131330230330 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3233203303111321-3231322112033101-2221200310110223-1202313122212102-1223331231202221-2121122131101331-1010313202012022-2203102032221023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212020031321011-2132323223303322-1302120312020101-0103031031132021-1302131302112033-1211222311112200-0001012221133121-2320310232103213"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 013233200310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-3313330222303111-0130121131122000-3133312212232331-3320033001102002-3103320033332233-3012003133000121-0101022022100302-3100303002002233"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable_path_normalize = {}
```

<a id="canonical-3001033102223120-3312130002210223-3333020232210330-1033313130002202-3221212021310322-0100131021021323-0322212231003123-1102031013022120"></a>

## Direct properties — enable_path_normalize / 013233200310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131303300131031-3013210232203201-1210012223311332-3123111210010221-1320312230213202-1013033021333201-0200020122200322-3223301213200110"></a>

## Next pages — enable_path_normalize / 013233200310 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110100312010130-2123212030300232-1023300320003213-2011330212320233-2300022031113202-1230323022223021-1333231230033322-0301101020313330"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — http_protocol_options / 312200322100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-1100331302233013-1321201200300220-3112100011031002-2121120332110332-1122310023001102-1100330103331332-2220030113231332-1323021230122122"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002321020211312-3310220121322113-1120202101203132-3130203320031321-3313132102130231-2122221203320022-0121032122100113-1310111220131011"></a>

## Direct properties — http_protocol_options / 312200322100 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-013.md#canonical-3001301323103222-2121231111201202-3033021201331202-2013203013231122-1022130313032203-3102302300331331-2322130230020110-1211323202320233): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-013.md#canonical-2233033312213312-3211033203222131-0330233230121212-1311330333020333-0230212230320330-1231010012113223-2230220213021211-0121110031221132): complete subsection reference.

<a id="canonical-2003011030002012-1021003012012313-3111210221323122-0002031321031102-1030321110303320-0000323110010302-0322330030010001-0131221103010322"></a>

## Next pages — http_protocol_options / 312200322100 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-013.md#canonical-3001301323103222-2121231111201202-3033021201331202-2013203013231122-1022130313032203-3102302300331331-2322130230020110-1211323202320233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-013.md#canonical-2233033312213312-3211033203222131-0330233230121212-1311330333020333-0230212230320330-1231010012113223-2230220213021211-0121110031221132)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031221311123130-3201011022003113-1003312123322100-0310132221112232-1122213132122110-1000220131223212-3101332222233301-0031010210013313"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 223110111312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0112211232330301-0132003111330112-0121020101031312-2301131220122310-1203223132033013-3301122310011133-2030123113103303-2030211200033030"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331113313313300-0002220332331131-3303203301221013-1231303000103313-3033012201312330-3312123313311332-2322302301020300-2003310001320013"></a>

## Direct properties — http_protocol_enable_v1_only / 223110111312 / 3

- [header_transformation](resources--workload--reference--group-012.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203): complete subsection reference.

<a id="canonical-0011012022303312-2320011121023221-0011131222100233-0033122310023221-1203220110201012-0013001003131322-3332202021323313-1321210102301002"></a>

## Next pages — http_protocol_enable_v1_only / 223110111312 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223032131002223-3233301103201301-1231313122121122-2200123232323011-3202012232231111-1303232232320302-2010111032133121-0201320321212333"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 100032130220 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1213001220033223-1222233021101330-2010103112302322-3233302001000300-3212100230133122-0102023213113033-1103332311320333-1033102002230310"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033020010013001-2100023322210233-2210011130201212-0222222110312202-3000233203311123-1132323212133320-2322130013113022-1200110021011001"></a>

## Direct properties — header_transformation / 100032130220 / 3

- [default_header_transformation](resources--workload--reference--group-012.md#canonical-1231232201100230-3232031132200011-2030110220331233-0022223202311032-2332231110211321-0022003220123032-2012102022313233-1033100030111223): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-013.md#canonical-2013011031103030-1210030331231320-3122133311033210-2310331322212101-1000121001321020-2001332010221033-2232310222330223-1320311010132030): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-013.md#canonical-3132322022001003-1100210323013220-2022130011201130-3330331210103031-2202332123113302-3030130333133213-0131322320120121-2302011103131201): complete subsection reference.

<a id="canonical-1230312100333010-1312012322132311-0010321220113102-3312031032300000-3321221133102302-2002021101200203-0231030323132203-1003230022033133"></a>

## Next pages — header_transformation / 100032130220 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-012.md#canonical-1231232201100230-3232031132200011-2030110220331233-0022223202311032-2332231110211321-0022003220123032-2012102022313233-1033100030111223)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-013.md#canonical-2013011031103030-1210030331231320-3122133311033210-2310331322212101-1000121001321020-2001332010221033-2232310222330223-1320311010132030)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-013.md#canonical-3132322022001003-1100210323013220-2022130011201130-3330331210103031-2202332123113302-3030130333133213-0131322320120121-2302011103131201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1231232201100230-3232031132200011-2030110220331233-0022223202311032-2332231110211321-0022003220123032-2012102022313233-1033100030111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
