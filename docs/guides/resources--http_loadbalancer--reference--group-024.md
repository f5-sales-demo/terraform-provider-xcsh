---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3102220322111132-1110113122103032-3331110333201233-1202001133000122-0113331333221233-1310320303312203-0200201231102322-1212101312233130"></a>

## duration property — hours / 130130330013 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-1101103333032313-3232133203211030-1322333202220132-2221103003330320-2132320211123211-1103022203303103-2322232201030023-2020021123100032"></a>

## Next pages — hours / 130130330013 / 5

- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-023.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1123220322033011-2022131300031302-0331123002023223-0013333220111023-2103032022232203-1331212310031113-3222233201131302-1132332303222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303220032301032-3333133000213011-0102220231223113-3031331312101321-0331112202012123-1012202313320133-2320333103331331-0130332022132211"></a>

## rate_limit.rate_limiter.action_block.minutes — minutes / 332012231301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-023.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-2203203013223101-2200113021201011-1100330032330331-0103212031203031-2232031111002220-1110121103321222-1131312301030131-3313132313113021"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101112033323202-1330030303033010-3222233002330022-1132303323003122-1302133301030120-3231330320202222-2203121320331312-3022113112002003"></a>

## Direct properties — minutes / 332012231301 / 3

<a id="canonical-0101133120101020-1322213102101030-2303033022031013-3102021121101130-2230330302113210-3301003023203132-0230111310213212-2310230212012133"></a>

<a id="canonical-1111032011211132-3000121233220313-0022223330130200-0332312120110220-2231310120311030-3320123321100132-2230010012032103-0333101212113130"></a>

## duration property — minutes / 332012231301 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-0011303210200113-1030231113221133-3012100130000212-2223232020103233-2232312222032220-2230112201222300-3030131313223222-0013302002222223"></a>

## Next pages — minutes / 332012231301 / 5

- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-023.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100321133312232-1010213301300133-1221233020303230-0320102001333003-3133303321011103-3133132232103001-0303113322032210-0221022013233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102012013100311-2103313203303011-0212002033022233-0033203223320022-3000121332032202-0101101000132100-0230210300311131-1312222133010011"></a>

## rate_limit.rate_limiter.action_block.seconds — seconds / 133312222310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-023.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-2102331331230211-3312013003332002-0332203201212003-2310331130213023-1121022213321120-3223001011213330-0331312030203110-3032321301032322"></a>

Type: `"object"`. single nested block, Optional.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113302301010220-1023033323232012-3323231130321133-2113330002123221-0332011012331220-3132100013010210-0100220132222001-3303303132233332"></a>

## Direct properties — seconds / 133312222310 / 3

<a id="canonical-2033332002002300-2312210033323220-1311320033133103-2012022131322322-3303032331103121-2230301133121322-2312302012033331-0310030120210221"></a>

<a id="canonical-1120022321200212-3131102122322100-1100002000010020-2310321000323210-1220102133203000-3313333222121120-1121122323031303-0132212202130212"></a>

## duration property — seconds / 133312222310 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-2301221011130100-2032300011221011-2032333030203012-2123223013311300-0102320120231023-3030121120311301-1010203201133012-2112322010000323"></a>

## Next pages — seconds / 133312222310 / 5

- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--reference--group-023.md#canonical-3103202121112302-1221322032222100-1212030133113120-2323120022200233-3321232113311332-1033310032011022-1300011312223023-0333001032310333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3232033201121333-3023321102300202-0203022310220000-3131212232330322-3113303131212013-0200123021011021-1233230122113213-3103023100232011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233123221120223-1331203301311113-0021222121332322-1023102222222310-3310222213020220-0100032211221021-0133110320330133-3023113131113323"></a>

## rate_limit.rate_limiter.disabled — disabled / 113233101301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.disabled

<a id="canonical-0010013332033122-0111110320233332-2112113320311311-1230012102202313-0201013312020323-3030120221312133-3120300303002213-2113132201032112"></a>

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
disabled = {}
```

<a id="canonical-2101202010112022-0031232320210310-1012123033012231-1130003310102332-1020322222302222-3223131022311011-1320020011013231-1221231110301122"></a>

## Direct properties — disabled / 113233101301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110310023133223-2000320200331013-1102121311212022-2210312001200210-0121303203133210-0210233002002030-2002133012011022-0222300223321100"></a>

## Next pages — disabled / 113233101301 / 4

- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102111021333002-0123232232020100-2131332023332322-3313212002110011-2232211323010132-2311001313201023-0331110230013312-1100303213132202"></a>

## rate_limit.rate_limiter.leaky_bucket — leaky_bucket / 223120223321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-2131030311121003-1030133220332212-1011212013002131-3231113221020320-2321223231122300-3301330201310330-0202201110222010-3123222021301022"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

<a id="canonical-3213013200002331-1130023002211130-2211213101303031-3310332012221102-2013131032033320-0012000333213300-2100213100212023-2012110001212220"></a>

## Direct properties — leaky_bucket / 223120223321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211222222201002-0032320021231232-3230320013332002-0110021001001302-1122311220003100-1200133331303103-2220003302003132-1033311301111113"></a>

## Next pages — leaky_bucket / 223120223321 / 4

- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303221200330202-3123321021000110-2332300002133112-2220222003033013-2013103231112210-3132200122112030-0221300021323023-2013233123003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311231312211001-0201111322233100-0302301002103033-0203000220301120-2021211111323201-2230113013120121-3302332123311210-0301331313001320"></a>

## rate_limit.rate_limiter.token_bucket — token_bucket / 233033102332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-0100203310120213-3200101121020211-2202323212003001-1011323011030332-1100301103012221-0112021202023321-3112223113023112-1111200023011110"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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
token_bucket = {}
```

<a id="canonical-0313221311110210-2123110122221110-1120311303011012-3330212312333332-2313122300112312-2030132010012303-3210131302030321-0021022233213233"></a>

## Direct properties — token_bucket / 233033102332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211202012331221-2200232310320022-2222131033132003-3020233321312231-1332112211331130-0311223213112012-0120333212310031-2033212201332210"></a>

## Next pages — token_bucket / 233033102332 / 4

- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-023.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302212102131332-2013333122123320-0332022333133022-3311100202303201-1001113302300230-2322333313321133-0233300311131302-1003021301303210"></a>

## ring_hash — ring_hash / 323223230023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- ring_hash

<a id="canonical-3013210122112121-0122221322331320-0013232112123331-2012132320331222-1132231302310032-2101233212003331-1003333001210203-3312200031323332"></a>

Type: `"object"`. single nested block, Optional.

Hash Policy List. List of hash policy rules.

Upstream description:

List of hash policy rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_policy")}
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
ring_hash {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332100331030320-0212133031000123-3333001311132310-1300112220102030-3032303302131030-3111013003322132-0300300230202223-0321031033323303"></a>

## Direct properties — ring_hash / 323223230023 / 3

- [hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313): complete subsection reference.

<a id="canonical-3013113100013012-0130213111233300-3212300013312111-1301313312222103-1233122131010132-2311203200103000-2112222310210330-2121332023010023"></a>

## Next pages — ring_hash / 323223230023 / 4

- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223223003010302-0230300113113120-0011000323022103-3213302113210330-1313302022021321-0211203100323110-0233032313220301-3100022003013133"></a>

## ring_hash.hash_policy — hash_policy / 010002110121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- ring_hash.hash_policy

<a id="canonical-1001233103230101-1000300000232320-2033102230313003-1310230311210113-1012230230203331-3223231202203232-0011132203210031-0130133230011322"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Upstream description:

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "header_name"),
  validators.ConflictingListObjectAttributes("cookie",
    "source_ip"),
  validators.ConflictingListObjectAttributes("header_name",
    "source_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113203213122220-0010200033110212-1010212120100230-3221032333233012-0000000200202220-0210303302221111-0010332231203232-3012112021002231"></a>

## Direct properties — hash_policy / 010002110121 / 3

- [cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022): complete subsection reference.

<a id="canonical-3130221302202133-2230001003300020-0010021133302231-1232033202332221-1010101011032233-3013000312011331-0301123102131101-0221313130020111"></a>

<a id="canonical-3213233031012232-2102220323312001-3212332120223031-2200323312131331-1303101131100323-1323013020012100-1132222132322213-1213130211032333"></a>

## header_name property — hash_policy / 010002110121 / 4

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Upstream description:

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022230202300222-1103201130121133-2020232311113203-0010032322013233-2231112201012302-2020302313220222-3122111300011000-0033301111001310"></a>

<a id="canonical-0112010311012000-2302310021211301-0232012103313010-0322213033302202-0112321212010002-1122022202030023-3300101230032012-1000301013130213"></a>

## source_ip property — hash_policy / 010002110121 / 5

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Upstream description:

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="canonical-3100100313301221-2003203103011131-0210322222110321-3310323303021000-0013331320110102-3021311233100030-3213102021100013-1211321323300001"></a>

<a id="canonical-0333212312303320-2101021312032312-1312000031000201-1320301330222312-2131331213311200-2323101002111032-1222023321121233-1223023330210011"></a>

## terminal property — hash_policy / 010002110121 / 6

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

Upstream description:

Specify if its a terminal policy.

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

<a id="canonical-2020030122023230-2112231211122121-0112113311101032-1203231031213003-2301212322033003-3231322023001311-0120111031210202-2201111223223102"></a>

## Next pages — hash_policy / 010002110121 / 7

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101123332113122-2333210302230022-0321130102231013-2103030331101000-1121113031002210-3213031120132201-2301202031223233-3302011331031201"></a>

## ring_hash.hash_policy.cookie — cookie / 230101223112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- ring_hash.hash_policy.cookie

<a id="canonical-1123221300031301-0233102300131322-0231223220233120-3303232200331313-0212311210210200-3311321223022320-0020111333303303-1231013112310213"></a>

Type: `"object"`. single nested block, Optional.

Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and
hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first
request from the client in its response to the client, based on the endpoint the request gets..

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
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
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313000002211333-3032122211233312-2112223020120322-2100331102013131-3122231110012232-2013100230111220-2020013323333023-0012332030212322"></a>

## Direct properties — cookie / 230101223112 / 3

- [add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100): complete subsection reference.

<a id="canonical-2112003132221332-1102011332030131-3301333332120300-3001301302323033-1211010120332021-2323323021010222-3023033010103221-3222112232231223"></a>

<a id="canonical-1200322331310201-0213111302331000-1123030330221101-0310222012032113-2032033121022012-3212200011133020-1031111023211120-0033201021121201"></a>

## name property — cookie / 230101223112 / 4

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0312330232332012-0031000233000213-1120331101023313-2302110201213031-1313012103322132-3201222312320331-3012323110133301-1333332032130123"></a>

<a id="canonical-0323003211220232-1330033011112133-2133121310131020-2102000331023031-3221021231132230-0023021202300300-0131300301323102-1102210220210221"></a>

## path property — cookie / 230101223112 / 5

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301): complete subsection reference.

<a id="canonical-1132130012331100-0310322232232230-2022212013030023-1232113320022300-0100303231132003-1322201301120203-1201022110222211-0023001001103103"></a>

<a id="canonical-1110102320222133-0031130302212303-0310233011110102-0320202111100220-1210111110023301-1203330333330002-2311212302111210-0013312101113321"></a>

## TTL property — cookie / 230101223112 / 6

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

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

<a id="canonical-3231130123220020-1132030011220030-1221123232300303-2013002010123131-0122311213101221-1003220211122123-2102302200012330-3001002333130200"></a>

## Next pages — cookie / 230101223112 / 7

- [ring_hash.hash_policy.cookie.add_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031)
- [ring_hash.hash_policy.cookie.add_secure](resources--http_loadbalancer--reference--group-024.md#canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212)
- [ring_hash.hash_policy.cookie.ignore_httponly](resources--http_loadbalancer--reference--group-024.md#canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203)
- [ring_hash.hash_policy.cookie.ignore_samesite](resources--http_loadbalancer--reference--group-024.md#canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213)
- [ring_hash.hash_policy.cookie.ignore_secure](resources--http_loadbalancer--reference--group-024.md#canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100)
- [ring_hash.hash_policy.cookie.samesite_lax](resources--http_loadbalancer--reference--group-024.md#canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230)
- [ring_hash.hash_policy.cookie.samesite_none](resources--http_loadbalancer--reference--group-024.md#canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313)
- [ring_hash.hash_policy.cookie.samesite_strict](resources--http_loadbalancer--reference--group-024.md#canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0233023321022032-3211013303101231-2222233200031322-2133030222102003-2302321230212320-0333010103001021-0120230300323003-3022321112010031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031203302003310-1302002011303223-0303102333233131-0333222023022130-3001133312331123-0103122313112003-2022121311212110-2002312032211332"></a>

## ring_hash.hash_policy.cookie.add_httponly — add_httponly / 331010121112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_httponly

<a id="canonical-3133310313000111-2031021211113022-2333121132001200-0002000320113033-1032321122023211-1203222021320021-3030013320002321-0023332000013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-0220323132320232-3000133303030302-0320021000031031-0111031130303010-0212031223003232-3121121301112302-3111233001000012-3001103203303122"></a>

## Direct properties — add_httponly / 331010121112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200210030213030-1213122111113021-0112322112300123-2122302120220332-2200201330221012-1002021303031002-2211203000001012-2300000313322113"></a>

## Next pages — add_httponly / 331010121112 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3211223231000113-0222213231133230-3212131130101231-1001311011131331-0133032202322231-2032323232301021-0111130223012020-2300121033030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330202310220022-2132233212121131-3033313031232131-3303113232010013-0301233120102000-2123210022322112-1120232212331031-3101002312101321"></a>

## ring_hash.hash_policy.cookie.add_secure — add_secure / 323210122103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.add_secure

<a id="canonical-2223202333002113-3121121013230322-3100033102002300-0303013132133101-2002130031223011-3223113000330122-2330300202320002-3302130321002233"></a>

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
add_secure = {}
```

<a id="canonical-0323203220212120-0002310123132312-3313301203223100-1313320121232220-1020000001220220-2023233021333223-0201033002310113-0030320001213121"></a>

## Direct properties — add_secure / 323210122103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121033211030332-3302012322001123-1220321313310111-1221323200220302-0200221022301333-3322320002220110-0012322031231000-2331121310110202"></a>

## Next pages — add_secure / 323210122103 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3300213232032231-2313023333310112-1121011103222021-0302330321101021-2013003213013331-0213002111110013-3101131111310201-0200002112231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333103330222021-3012311222102231-1321113331210231-2111301330321212-0230132133300222-0002032231133230-2203031003011322-3311001200302121"></a>

## ring_hash.hash_policy.cookie.ignore_httponly — ignore_httponly / 030030003103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_httponly

<a id="canonical-0102320010131233-1133121101111033-2311002213233001-1103001232310101-0220101200020010-0220321110212200-3301121111230021-2323122121003102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-0221102230333020-1300011311303212-0021000231013310-1220201200300130-2010130003310303-2210122211311032-2332200323010223-2123130212001331"></a>

## Direct properties — ignore_httponly / 030030003103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302022012013122-3331002322330220-2021311011303222-3303102221020221-2113210123101123-3221100311231000-3201210030130011-3301001132002230"></a>

## Next pages — ignore_httponly / 030030003103 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2332333230022233-0132022121301303-0103202022312020-2313313001110002-0010110100130212-2310031310302213-2102022001003100-3103303321201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230130123330101-0222213331210011-3222310121222223-3130001013120313-2221122101133222-2322020032321301-0302112022230123-1311002312221131"></a>

## ring_hash.hash_policy.cookie.ignore_samesite — ignore_samesite / 203003232223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_samesite

<a id="canonical-3331020003223201-0012120213232031-1131223211002332-0112201012023132-0310021003100100-3113013012100203-0132032311221020-2130322123013200"></a>

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
ignore_samesite = {}
```

<a id="canonical-0011113231332323-1202221212313300-2131223302012211-1233022230333033-2102103322301112-2000132301330001-1233310102133233-3102202001210020"></a>

## Direct properties — ignore_samesite / 203003232223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302001200002002-2100332330022202-1011021131230223-1033320230112002-3002101030330230-1310131000123300-0020132003123020-0121101323322331"></a>

## Next pages — ignore_samesite / 203003232223 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1120222021213231-3003220120213132-1322210223123032-2300111330030321-2133103202003233-0211133213032211-1020320331332121-2112122222011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022003221231210-0200233203033300-1311032321001321-3331132223112321-1003121120023332-3222032132312101-3000220220330320-2222020300130222"></a>

## ring_hash.hash_policy.cookie.ignore_secure — ignore_secure / 213000333210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.ignore_secure

<a id="canonical-0333113130311222-0210223021200112-1213131031223333-1020320212103020-2133133332201312-1032021330020310-3322100020311310-1223302032102331"></a>

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
ignore_secure = {}
```

<a id="canonical-2222103220122330-3201000131201211-1203133210012302-2331202302313123-1102211132210131-0032312300102012-2221311202032023-2102031312200310"></a>

## Direct properties — ignore_secure / 213000333210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111102211201011-1110322000310020-0000201011001312-0122000021013220-0130003123330321-2313232302131131-3210202033333033-0133230001313021"></a>

## Next pages — ignore_secure / 213000333210 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1000320300010321-3113212121101033-2123003330212330-0323233012330112-0110012230030333-1210122230222230-2303323122231003-2120223222320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300232303002312-3103332212123133-2331221300321322-3213331100012223-2130121131302032-0000313222232220-0231103210033110-1030100110313211"></a>

## ring_hash.hash_policy.cookie.samesite_lax — samesite_lax / 000001202302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_lax

<a id="canonical-3322132213203201-1231030122102312-3313323023333302-1103130202122202-2321000011313322-1023332010301112-0003012211311132-2021101230311332"></a>

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
samesite_lax = {}
```

<a id="canonical-1320232210223012-0112210333003301-1112121000110031-2002310332322332-1033012311232032-2230213231103312-1002133202222133-3131310311230013"></a>

## Direct properties — samesite_lax / 000001202302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212032122232112-2010321002022021-1103130122011030-3203013332202102-0330031310233123-0223110303122033-0021030001103100-0233011000130022"></a>

## Next pages — samesite_lax / 000001202302 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210321212023202-1313312121321320-2320323300302001-2013223223222312-1111113110132012-0313112102002201-1122300112133003-2112203110300213"></a>

## ring_hash.hash_policy.cookie.samesite_none — samesite_none / 132132002131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-3202023323212002-3032033213100112-1123312131202001-2131312331303001-0113230000100210-3001210211011302-0111113312301003-2322003301000330"></a>

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
samesite_none = {}
```

<a id="canonical-3233311102020321-1003103300103112-2233211333112132-2121332130133132-2221303003110020-0301310332232313-2300213101003203-2001222031303123"></a>

## Direct properties — samesite_none / 132132002131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111233102010130-3303113120013232-3231212131002233-3210130222011213-3113131321133103-2113033213011330-0132002110311213-3010120013332031"></a>

## Next pages — samesite_none / 132132002131 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032001021002313-3232303103102031-3010302123120022-1313322203313100-1300121231220300-2221323203033110-3320113211003111-3333121001111221"></a>

## ring_hash.hash_policy.cookie.samesite_strict — samesite_strict / 313023313121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-2221232320232210-3102000210122201-0013322231302022-1213121230230323-0210210232213131-3323121133012011-3022332100130213-3032122111300113"></a>

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
samesite_strict = {}
```

<a id="canonical-3310201301000011-0223130132010122-1120133231001010-3131303000311001-1323223011332100-0201132033100123-3303303100132010-0221230022200230"></a>

## Direct properties — samesite_strict / 313023313121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031112211013023-2211213320332230-2013313221113323-2330131320111130-3122020221100330-0111320003321221-2102132330010100-1300211012011020"></a>

## Next pages — samesite_strict / 313023313121 / 4

- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3120201302000120-1101032023332332-1203301030103211-2301112131023100-3323333313021213-2102103120123013-3211232030113002-1310102202033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132103212103130-2123303211223331-0102132112323333-1213033002210310-1223010220232213-3123122321321131-2010111000000133-0130212231103323"></a>

## round_robin — round_robin / 033031022201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- round_robin

<a id="canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
round_robin = {}
```

<a id="canonical-0032202013022023-2032100303300123-1202120122022330-3230031031210110-0323213130332131-3112102102220132-1312123010102232-1133012022310030"></a>

## Direct properties — round_robin / 033031022201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020320213112013-3202202321222023-0302000013101200-3300013000022310-1230021320200203-2013112333123032-0123201133211030-2113113211112031"></a>

## Next pages — round_robin / 033031022201 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302333330103332-3000231123012200-1223313032013033-3110312130233300-3201000202121113-1221011200232010-0313220022331010-2221323130111313"></a>

## routes — routes / 000001003111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- routes

<a id="canonical-1022112302210201-1010011312111323-0033100111313112-0212233102202102-1021022230231123-1121103333323120-1231212033301110-1123111133332122"></a>

Type: `"object"`. list nested block, Optional.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Upstream description:

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("route_state_disabled",
    "route_state_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020332323331332-0013300321000110-1101323220321130-0201331131231111-1333323212313323-3203111030201212-3222303202132212-3010030323231312"></a>

## Direct properties — routes / 000001003111 / 3

- [custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130): complete subsection reference.

- [direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311): complete subsection reference.

- [redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223): complete subsection reference.

- [route_state_disabled](resources--http_loadbalancer--reference--group-024.md#canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303): complete subsection reference.

- [route_state_enabled](resources--http_loadbalancer--reference--group-024.md#canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203): complete subsection reference.

- [simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232): complete subsection reference.

<a id="canonical-1010200102203001-2033002321011102-2120023202021023-0122210333333231-3303233012332302-1011132020001122-1332211201220222-1102301200020203"></a>

## Next pages — routes / 000001003111 / 4

- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.route_state_disabled](resources--http_loadbalancer--reference--group-024.md#canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303)
- [routes.route_state_enabled](resources--http_loadbalancer--reference--group-024.md#canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310333203312222-2032203322200333-3103123323230211-1320102023312023-0013322211322022-0302203320221221-1122002313223023-1000303311113320"></a>

## routes.custom_route_object — custom_route_object / 332132122132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.custom_route_object

<a id="canonical-2132120132332033-0111230301202203-1332310120131330-3101230211330320-1302310110310033-2231032202200231-0302213013300201-0020203132023302"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122011203302132-2210201030021322-2023211301013312-2133220003332101-2122113120131121-1212231330020012-0101003103110023-2013113230030220"></a>

## Direct properties — custom_route_object / 332132122132 / 3

- [caching_disable](resources--http_loadbalancer--reference--group-024.md#canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-024.md#canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110): complete subsection reference.

- [route_ref](resources--http_loadbalancer--reference--group-024.md#canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223): complete subsection reference.

<a id="canonical-2030311221102101-0132302000302202-1001003013302030-3231222221232220-1223020302302312-0001022203212202-3222222330131013-2011010320323030"></a>

## Next pages — custom_route_object / 332132122132 / 4

- [routes.custom_route_object.caching_disable](resources--http_loadbalancer--reference--group-024.md#canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312)
- [routes.custom_route_object.caching_inherit](resources--http_loadbalancer--reference--group-024.md#canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110)
- [routes.custom_route_object.route_ref](resources--http_loadbalancer--reference--group-024.md#canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333112221002200-3211331201000132-1121320120003010-2201001333231012-3320001133020110-2322300211012230-1012212202023333-3130223023003212"></a>

## routes.custom_route_object.caching_disable — caching_disable / 003030102220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_disable

<a id="canonical-1300132233003321-2310012321003102-3313000230313012-0113022222000330-3212023132023010-0023031201321130-3322202333303202-1100023203132133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

<a id="canonical-1013221301113331-0230232131321032-2123232223311111-2321312233012110-0032220013020220-3022231110023101-2203230221032010-0120133220103000"></a>

## Direct properties — caching_disable / 003030102220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032210022020020-3013320113212121-0003210221231121-2300333321133000-3113013101132320-3213333211311231-3132002022110200-2200033313111110"></a>

## Next pages — caching_disable / 003030102220 / 4

- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111333213231320-0031102302001002-1303312012312210-1003021021302221-3222033131211200-0320321132011311-0332321123113130-2201010121321112"></a>

## routes.custom_route_object.caching_inherit — caching_inherit / 233010311210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_inherit

<a id="canonical-0023203021031101-3102012330230032-2322300303011202-1102031110221332-2122231020331130-0013003023132321-3233331212131000-3111121233333103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

<a id="canonical-1331102203220102-2232113000313132-3221231303130002-2232220311003321-2313220203313333-0313110301312110-0001123330331133-3312322003323321"></a>

## Direct properties — caching_inherit / 233010311210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322101211333000-2223332233202032-2030203132002231-2030212101032102-0123233002133001-2032111220201131-1323023010220230-0301102121202010"></a>

## Next pages — caching_inherit / 233010311210 / 4

- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102303303101220-2220313232201222-2311212111013010-1023203313233111-0333303322101003-3011203321211000-0022233012120313-3313232320032320"></a>

## routes.custom_route_object.route_ref — route_ref / 301302122012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.route_ref

<a id="canonical-1123021100312002-1020112010103001-2313033300010013-1303122122012223-0322203113312222-3001003121332000-0101330122123222-0033201313203300"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311321212322103-3101311301011101-2031320300213121-3003320022201310-1210130033302301-3202212201010300-2321101213220000-1222103300233331"></a>

## Direct properties — route_ref / 301302122012 / 3

<a id="canonical-2323000200333123-0132121222212333-2210330112033001-0112112303232223-1213013021101003-1213132222113321-2330222001330012-2320213231123322"></a>

<a id="canonical-3103132333232223-0210101111301211-2313102222332013-0112132211230132-3103123031231303-2321313013123321-2002130113113013-0022123222000133"></a>

## name property — route_ref / 301302122012 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2201320121001121-0103233133130321-2101212203012032-2033101021032220-1230300223120133-1111112032301030-3312332122301100-0220322033203122"></a>

<a id="canonical-2121010103332203-2032132010120110-1230202312202300-2332331200212222-3330011233332033-2320031323102303-3023020131230301-0030031132112002"></a>

## namespace property — route_ref / 301302122012 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0221030111013133-1101230221300312-1320131302031110-3310103212313113-0130122232133121-3313211110000220-0120311133031233-3233210101331211"></a>

<a id="canonical-3332133203323123-1300323020021123-3300011233232322-2312300312321202-1212202203013013-3130133001211202-3120332021210233-2132031201030032"></a>

## tenant property — route_ref / 301302122012 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1331221022231030-0203200011222032-2300210032310102-2231200201102302-0220323003202122-1320211211200310-3113302012321302-3201131331321002"></a>

## Next pages — route_ref / 301302122012 / 7

- [routes.custom_route_object](resources--http_loadbalancer--reference--group-024.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231203213112001-1101032030331021-1020030011212321-1300000102330231-3010001302223211-0111013003103303-2211200202013321-3211210323332211"></a>

## routes.direct_response_route — direct_response_route / 032213333220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.direct_response_route

<a id="canonical-2212332230001201-2200303013102020-2331330332220323-1212131312310033-0303100201123223-3332120313321120-3323111102000031-3120322323221310"></a>

Type: `"object"`. single nested block, Optional.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

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
direct_response_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303333213133003-2121001302000121-3330230011310322-1130010131201222-3120003332332031-0100023121321202-2320310321322300-2033202022113323"></a>

## Direct properties — direct_response_route / 032213333220 / 3

- [headers](resources--http_loadbalancer--reference--group-024.md#canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133): complete subsection reference.

<a id="canonical-1122320301230101-1201300111222321-3200131200000200-0011323233232012-2001002210321013-1223010111113202-3302221330003001-2012320102223311"></a>

<a id="canonical-3232312110313302-2120210110320301-2330032020023110-0133020020123122-3320033132213111-2313033102031031-3100222001320010-1100113201331002"></a>

## http_method property — direct_response_route / 032213333220 / 4

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

- [incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-024.md#canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233): complete subsection reference.

- [route_direct_response](resources--http_loadbalancer--reference--group-024.md#canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320): complete subsection reference.

<a id="canonical-1123102021231000-0220222120000312-1130110112001323-2313130000210223-2123111310131211-0320002232100323-0230330122001110-2001100122001100"></a>

## Next pages — direct_response_route / 032213333220 / 5

- [routes.direct_response_route.headers](resources--http_loadbalancer--reference--group-024.md#canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133)
- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221)
- [routes.direct_response_route.path](resources--http_loadbalancer--reference--group-024.md#canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233)
- [routes.direct_response_route.route_direct_response](resources--http_loadbalancer--reference--group-024.md#canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013322223213310-2021313020321210-2012032031210100-0113001031130131-3311010233312212-3303322213331130-3312313030321202-1311100301303133"></a>

## routes.direct_response_route.headers — headers / 210123022023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.headers

<a id="canonical-0012332300232233-2300123121101322-0230222021131133-1331332110303131-2000331133321110-0120322233122301-0033120130031021-0332222011302133"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0122021221020332-1111130033232312-1330001012102123-0000331013220202-0010103223113133-2311233121002322-1222300001031200-0033222011321002"></a>

## Direct properties — headers / 210123022023 / 3

<a id="canonical-2013030023011233-0003132110221320-0233300111103123-2213132100123213-0303220033321100-0322200212300331-1133002002203333-1202033030221101"></a>

<a id="canonical-0101213020110003-2131331222302130-2133330010203210-3303030113000220-3013322233012230-2212233032312332-0231031133102113-2123110232211032"></a>

## exact property — headers / 210123022023 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0030210001022320-2130302220112322-1133000102110022-2231303320003321-3132310003201332-0023001033223110-2331233332121103-0303221031322323"></a>

<a id="canonical-2120322223313323-1320313303311203-1321300201330110-3211021001212031-1203320331200011-2121010332222201-2032102111013101-0213202201332111"></a>

## invert_match property — headers / 210123022023 / 5

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

<a id="canonical-3013220100011213-3203332322113000-2021133221211121-0123301033332013-3121000110102013-0022330003303012-3321112022031132-0220231123103000"></a>

<a id="canonical-3211222113013123-2013101001321232-1311231231031011-1320301032132113-1302303211330201-3202312000300110-2023211111032123-3201313013213010"></a>

## name property — headers / 210123022023 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1133322321301032-2133311133200113-0200200133231332-1112230202233003-1031302310303012-3030123200330211-2112312022010103-1310202030232213"></a>

<a id="canonical-1331330131000222-0320011323201311-3110212131131100-2313003013130030-3311021100302132-2313301123333122-0020331031131013-1112312213120131"></a>

## presence property — headers / 210123022023 / 7

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

<a id="canonical-1123212112231001-1301302003301210-2211323232003121-3031231333030220-2101023012222322-0130122112101310-3103031133112002-0200131221312123"></a>

<a id="canonical-3001100131001032-1301323120122013-2203221133120202-3232111302232010-1122033132132210-2120231100133121-2312120103221012-3031300200302300"></a>

## regular expression property — headers / 210123022023 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2221123032310303-3221020331123213-3230302303301230-2023111130300032-3033210232223203-2121033312000011-3022002333310302-0032023032002223"></a>

## Next pages — headers / 210123022023 / 9

- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003122032113122-1121123110231012-2031232131133222-2210233100201232-2223103132133330-0120312101013322-1000102130033021-0133311212123012"></a>

## routes.direct_response_route.incoming_port — incoming_port / 232221303300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.incoming_port

<a id="canonical-0010030313231201-0201232322323211-0220321212230231-3011233202110100-3033311203002302-0303131010110200-2130333232032332-2130032130012302"></a>

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

<a id="canonical-1223223003321303-3212101021322213-0020132322001013-0011100321120111-0101123022031323-1221101101011313-0332112221122121-2031323002123233"></a>

## Direct properties — incoming_port / 232221303300 / 3

- [no_port_match](resources--http_loadbalancer--reference--group-024.md#canonical-1331102132201121-0032301222132323-1023203312102303-1220103202030130-2320013100232330-0232333120031311-3300203113320032-2113010021211202): complete subsection reference.

<a id="canonical-2310311221220313-3020200222211300-1102313020123210-0013232320300033-2312201321013310-2211230300001231-1330211300023002-2233221330002213"></a>

<a id="canonical-1132110121300121-2002211001132222-2021023210322111-2120120032313012-3301230320130232-1202211110010323-2301031103123101-0212232013122113"></a>

## port property — incoming_port / 232221303300 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3022020013313131-1201213001231210-1302003100120331-1122113120300221-3302211012230010-3032102211020311-1030120101223100-2303312232022033"></a>

<a id="canonical-2232031320011111-1011333221010301-0030012012033201-2122220020100323-2302223202202210-2203120012311331-2330220231121330-1133100322200300"></a>

## port_ranges property — incoming_port / 232221303300 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0222210110032322-1132332210313113-0131221223032022-0232211301023131-3331302031103332-1130131131330213-1000220120200023-2131332200223133"></a>

## Next pages — incoming_port / 232221303300 / 6

- [routes.direct_response_route.incoming_port.no_port_match](resources--http_loadbalancer--reference--group-024.md#canonical-1331102132201121-0032301222132323-1023203312102303-1220103202030130-2320013100232330-0232333120031311-3300203113320032-2113010021211202)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1331102132201121-0032301222132323-1023203312102303-1220103202030130-2320013100232330-0232333120031311-3300203113320032-2113010021211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231133230322300-3233311003200023-2232132002203320-0013110323322310-3301223301013220-3001113002030031-3323113230213101-2202113331000113"></a>

## routes.direct_response_route.incoming_port.no_port_match — no_port_match / 211330011021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221)
- routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-1120100133113130-0230002230001113-3032202022101320-0300101221113200-1123113203302110-3231320202013320-2330113103213220-1321133332012022"></a>

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

<a id="canonical-3033010321020320-3302322330302101-1021122201121211-1003123202331332-1113132101132121-1032212323223221-0331123030022321-1232200022320233"></a>

## Direct properties — no_port_match / 211330011021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033112111300032-1031313233230222-0303202301332132-2321122020103321-3111103112203113-0033022201211020-3320003213023203-0021323320323112"></a>

## Next pages — no_port_match / 211330011021 / 4

- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310132333010120-1323213310011123-2002321121022032-0023322020030010-3132202302102131-3233311311111223-3133221131323330-2232332031210332"></a>

## routes.direct_response_route.path — path / 320233111201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.path

<a id="canonical-2123030001220310-0032112112231312-1323221203002000-1333210210232011-3032101112112232-0330203300123220-0310221000213310-1323001232013132"></a>

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

<a id="canonical-2231332102202001-1100111202201301-0302310012303321-0222200111132101-3001001131303020-2300123103203221-2013311221100230-0013103122013301"></a>

## Direct properties — path / 320233111201 / 3

<a id="canonical-1210112120201033-0003132032113310-3330302021130333-3023000200111020-0323130123132213-2012301112222202-1130201023302312-1223223200130210"></a>

<a id="canonical-3000312330103302-0202023220002212-3012322110132213-3020103310000312-3021201032112001-1212331231333310-3333321133332203-1001230102100330"></a>

## path property — path / 320233111201 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3203200023320320-3022231201103301-2031322231130112-0032101323120120-2213311112211202-2003233011001002-1232220203213221-1221221010000230"></a>

<a id="canonical-0330021133102310-3323331022020223-3223032020100131-0133033132332233-3233033123000113-2230300113013202-1220210123332213-3331001113031122"></a>

## prefix property — path / 320233111201 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2113310112230021-2032110332210233-0311321031322301-3021232120221003-0210023230020023-2301231332100230-2131022113131331-0303133222311220"></a>

<a id="canonical-2223033113002333-3312131131331011-2013320331331123-1031012330121020-3001311132103110-3032233021121013-2222001122112201-0211330032333102"></a>

## regular expression property — path / 320233111201 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2100020312103121-1200331020022013-3112220111201031-0001310000003123-3031031203303011-2312332310131003-2201100010220111-0303020203122321"></a>

## Next pages — path / 320233111201 / 7

- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132133112100133-2201100022203202-3331023133003231-3232131131110013-1233302322223310-0100203113232201-0321021111002011-2300131211131121"></a>

## routes.direct_response_route.route_direct_response — route_direct_response / 010302333003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.route_direct_response

<a id="canonical-3213211133213302-0333001303300200-3003221012221322-2123301223023211-1023301101321300-0230210032320111-1132301103022012-0022032231023331"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301220030030031-0121221333012021-1233100233230113-1310313303313132-1023021313311312-1200322023320331-2100323031200012-2133113231113130"></a>

## Direct properties — route_direct_response / 010302333003 / 3

<a id="canonical-2121031131100330-2321212011300132-2023000001220011-0311001220122300-0303001021211010-2010210002220002-1130311211321311-1110223012303333"></a>

<a id="canonical-1132121123300031-0121210003011000-2321203020332102-2311332233133003-2301112030220112-2123123003013333-1333000123303202-2011222111100111"></a>

## response_body_encoded property — route_direct_response / 010302333003 / 4

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2301221330322221-2031020121230122-0030032113320211-3322101011131320-1330322111032223-2012132120231001-2210202103333123-0212222200103031"></a>

<a id="canonical-2022033111332330-3322312130103333-3113301210030011-0201110301312233-2031033213310233-1230211110003012-3123132120120331-3231212020131302"></a>

## response_code property — route_direct_response / 010302333003 / 5

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-2222223021033330-3302101320322312-0321032101212123-0002012310303122-0321001010301221-0021202132130210-2201011330222311-2031331001031102"></a>

## Next pages — route_direct_response / 010302333003 / 6

- [routes.direct_response_route](resources--http_loadbalancer--reference--group-024.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300123002010103-2323232011201201-2003210012102023-1122221203130102-0113110331133123-1310211232120102-1032132330220301-0222103210303110"></a>

## routes.redirect_route — redirect_route / 213232333300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.redirect_route

<a id="canonical-0311011002232332-1323022032033321-2322303332331220-2232311213123212-0102222022201230-2133233221120032-2010221020312131-3130020003032202"></a>

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

<a id="canonical-0220232021313312-2301021202123131-0000031130011022-2303120033223300-3310232012230000-2210231023020101-0323311220113120-1131210312212023"></a>

## Direct properties — redirect_route / 213232333300 / 3

- [headers](resources--http_loadbalancer--reference--group-024.md#canonical-1012101232011000-1123332000231212-2010303222022313-1111112132300001-3101031010322232-1321223232223321-2231031213011001-2202132202020132): complete subsection reference.

<a id="canonical-1330213120223223-0221132001132120-3302123321100321-0313011332203101-0233221232033022-3121322102330102-1202122122121012-0020220133103320"></a>

<a id="canonical-2201000303130321-2333031122003112-1202333123022022-2203300222203230-1131330203213022-3200103021202333-1201212221102121-0213232031330113"></a>

## http_method property — redirect_route / 213232333300 / 4

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

- [incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-024.md#canonical-2023020103020223-0112101233123130-1311203130032022-0133100102302010-0002330321210002-2113313131300313-1110223311232213-0310323102121021): complete subsection reference.

- [route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011): complete subsection reference.

<a id="canonical-0222121001212032-3311021012303110-1102233222100330-2313201323102121-2121133031300003-1022001102003322-1311101102020200-3300321212100212"></a>

## Next pages — redirect_route / 213232333300 / 5

- [routes.redirect_route.headers](resources--http_loadbalancer--reference--group-024.md#canonical-1012101232011000-1123332000231212-2010303222022313-1111112132300001-3101031010322232-1321223232223321-2231031213011001-2202132202020132)
- [routes.redirect_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101)
- [routes.redirect_route.path](resources--http_loadbalancer--reference--group-024.md#canonical-2023020103020223-0112101233123130-1311203130032022-0133100102302010-0002330321210002-2113313131300313-1110223311232213-0310323102121021)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1012101232011000-1123332000231212-2010303222022313-1111112132300001-3101031010322232-1321223232223321-2231031213011001-2202132202020132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300103223223023-2020123223130212-0330112302203303-1123213103131310-1030133010033233-0322233000323022-2103322302130220-2321333301301133"></a>

## routes.redirect_route.headers — headers / 321330300223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.headers

<a id="canonical-0001220321021231-2131102303311121-3330031003120230-3110003303320203-0132101112110031-0211333331131021-1103011332310131-3133111310211330"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1133111311021212-2001323231202330-0311100310323122-1023300201100012-1231100123120030-2102320323230321-1111120333110221-0211023233021222"></a>

## Direct properties — headers / 321330300223 / 3

<a id="canonical-0320131032231131-3221100221130023-2301002022033320-0023233121113222-0201333112312130-0102232000120211-3110222323222103-3013030223130102"></a>

<a id="canonical-0111031220332123-1103032201231102-2112323110300022-2201011232100001-0232312221102302-2200330302121222-2012301333211021-2121222301310021"></a>

## exact property — headers / 321330300223 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3030332321020301-3102320203103330-1013001030103212-0000133330332233-0010211230013130-0132333201311200-0131003210330201-1331030002302111"></a>

<a id="canonical-1131132233322112-1131123223121220-3001202023232032-3013331211201033-2121000230033230-2023033313310100-3330102332003320-0200222230001021"></a>

## invert_match property — headers / 321330300223 / 5

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

<a id="canonical-2132202002310202-0323230320130033-2112211000000032-1131323113122121-3120122131020130-1312311210110000-1120310030221233-3322103231112002"></a>

<a id="canonical-2331031211001012-3132302110320231-3213101203101120-1202013302023010-1320120102200100-3013203223323331-3033012132220000-3101312212101221"></a>

## name property — headers / 321330300223 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1322302131111011-0001013131300021-0322333130201232-3310313301321101-0222212303120302-2021333033110212-3023002131131313-2130310232223220"></a>

<a id="canonical-1200202010211302-1012200010223221-3233121132112031-1121300221232311-3031130232130112-1030221220230222-2033022001330122-3130020020012321"></a>

## presence property — headers / 321330300223 / 7

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

<a id="canonical-1302300231213323-2310023133103302-3102220220131013-2301233013103212-0300133313211313-2231131310020320-1131120331100001-2122000222203100"></a>

<a id="canonical-3022221220122322-1102121100311001-1030231002000010-0030331331210230-1022211213013120-2011122003102131-2311003022021103-2103101333231322"></a>

## regular expression property — headers / 321330300223 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3212312202111002-1032331133122020-2122323312332320-2231322232132332-0100112213122122-3212212221022132-0033112211123003-0010022210330223"></a>

## Next pages — headers / 321330300223 / 9

- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000320213312332-1033123223003232-0110013303223112-0100230201132223-1102232300113103-0123232123101320-2103133200000031-2201303133030311"></a>

## routes.redirect_route.incoming_port — incoming_port / 022221203032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.incoming_port

<a id="canonical-1130011123033212-3233313012011202-0210023031211022-3232122110203303-1001130102031321-1300323032023212-0120322121103003-1203212130310220"></a>

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

<a id="canonical-1101313131000122-2311200112001220-3100213220120203-1032313231101310-3203003211233302-0311233121113121-0200100013012113-1121222231012100"></a>

## Direct properties — incoming_port / 022221203032 / 3

- [no_port_match](resources--http_loadbalancer--reference--group-024.md#canonical-2101022032022113-0000033222103332-2120031101001221-0132231222332011-2022301023332230-0220203133310123-1200221332020011-1023312233230223): complete subsection reference.

<a id="canonical-2322220232233332-0200213123011321-2120201200211002-0222222332002212-0223223022303102-2321331101132003-2020320021120113-0213102121000102"></a>

<a id="canonical-3212312130121021-1301002321321200-0001301311100302-3012023122031031-2331310031313123-1033212100222001-0010232221311330-0303213120200302"></a>

## port property — incoming_port / 022221203032 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1033031321020310-0230103121032333-1231031011332332-1031013221330211-3222203332101133-0311023202000021-1101323030223330-2000123123211000"></a>

<a id="canonical-1333333023322133-3332201311333213-2011320300203112-1323303330202213-3022000231223303-2002322233121202-1020131032022011-1111223210331001"></a>

## port_ranges property — incoming_port / 022221203032 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2112311222133030-3132131322102230-0220011313232120-3020122120313120-1210201130111000-1100101002023112-3021112132001112-0000012331112012"></a>

## Next pages — incoming_port / 022221203032 / 6

- [routes.redirect_route.incoming_port.no_port_match](resources--http_loadbalancer--reference--group-024.md#canonical-2101022032022113-0000033222103332-2120031101001221-0132231222332011-2022301023332230-0220203133310123-1200221332020011-1023312233230223)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101022032022113-0000033222103332-2120031101001221-0132231222332011-2022301023332230-0220203133310123-1200221332020011-1023312233230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330230113332023-1101000303320000-0321120233021101-2013100312232303-1231110122322331-0000110033312033-2112322212312123-1022322033032121"></a>

## routes.redirect_route.incoming_port.no_port_match — no_port_match / 312110102010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101)
- routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0232321023233331-3022002131110110-3032233213031333-0113211111200323-0203001110201012-1101012003333200-3201001312033211-3330020212031210"></a>

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

<a id="canonical-2113021122231233-3021233010111122-0201303100010023-2112303120023101-1113221011020100-1003231021310133-3103232031133000-1200021310210300"></a>

## Direct properties — no_port_match / 312110102010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211300022303321-1100202222203023-2132332110210322-0112200122011322-1302131221203301-2232022022220123-3112322301203101-3131223232023213"></a>

## Next pages — no_port_match / 312110102010 / 4

- [routes.redirect_route.incoming_port](resources--http_loadbalancer--reference--group-024.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023020103020223-0112101233123130-1311203130032022-0133100102302010-0002330321210002-2113313131300313-1110223311232213-0310323102121021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113201032002123-0130213022003301-0301111221130000-0330231332313211-1210131022103112-2133020301133201-0312120033220220-0313131311330312"></a>

## routes.redirect_route.path — path / 031132312011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.path

<a id="canonical-0322332033000233-1023101112021013-0121233000230002-1211220012032102-2130203030130033-2332000313011330-2231131130020131-1132322312122321"></a>

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

<a id="canonical-3213222312020320-1222311031212202-2023010322201221-3310000011222333-1031110300321123-0001122001023122-2323323303203322-2312210310011320"></a>

## Direct properties — path / 031132312011 / 3

<a id="canonical-3312211100121222-1023203120303312-2112110013130231-3101230003333112-3333133103033033-0321210112010301-0333012302313201-1202200231011311"></a>

<a id="canonical-3130310210312130-2121031020222222-2201330311023132-3211022323332132-3311212030333030-0310011222133010-3013021301110303-1011211223131231"></a>

## path property — path / 031132312011 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1003033223102311-1023231112032030-2021000320231213-3321030301002301-1321022002311030-1222211313221030-0012322032113322-3211010220000001"></a>

<a id="canonical-3023023302133310-1222011333222011-0212322331120320-2202023112023023-1003102301113222-2233012130031220-3321113210333320-0003131111211303"></a>

## prefix property — path / 031132312011 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3323000211323231-0130020103111211-3323210211001112-3020333010331122-0002113203113023-2010313310222001-1133203311223030-3301011021211021"></a>

<a id="canonical-0002312110200122-1133203121121013-2122333003223213-2020213002000330-2113332111000132-1303321220113030-2232103011230212-0133201123023330"></a>

## regular expression property — path / 031132312011 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3212123300311102-3021300022132031-1013310030032312-3223301022001333-1122123033230000-1030213302201121-3200222100100113-3101103322211213"></a>

## Next pages — path / 031132312011 / 7

- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022231230022023-3121121000000211-2321102221131101-3031001120333123-1212200132321302-3301220101303223-3102101012313211-2121032003323102"></a>

## routes.redirect_route.route_redirect — route_redirect / 200330313210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.route_redirect

<a id="canonical-2320211100222130-1302322000233133-3031100330010023-3331301300332310-1221102132021100-2023133010131232-0333121022000330-1100131312102212"></a>

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

<a id="canonical-3002013120231133-1123101232222113-1011332113002013-3120020332232032-3002000112233200-1103220333300010-3110111022330131-3230102032003303"></a>

## Direct properties — route_redirect / 200330313210 / 3

<a id="canonical-3120120112132302-2123113031300233-1101322202101303-0122132203131023-0200112203133233-2213332223132023-3321200211302131-2320323112323020"></a>

<a id="canonical-0022300231303230-3232300213011333-0131013303133033-1100131123033331-3012002022101222-1312302133210300-2223121030220020-2003233130202013"></a>

## host_redirect property — route_redirect / 200330313210 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331100133113301-1323312211331300-2230331002021033-2221222320131200-1113211313122301-2210022032033332-3233321331300223-3222232013213330"></a>

<a id="canonical-1033110222133211-2120220210332311-1112230110320211-3132221130333313-2002033331230321-2032330233121002-3121000330223200-2022132313011212"></a>

## path_redirect property — route_redirect / 200330313210 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2212003322032101-3323223302333102-0001220132002003-3321303022000113-1301213002012322-0002011133321100-0112023023202111-1021121030303312"></a>

<a id="canonical-3330203112101300-2312203122212030-2223110201331330-2101131222011212-2113310013102112-3302033113320231-0221010022302132-0132333122003102"></a>

## prefix_rewrite property — route_redirect / 200330313210 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3102311302013113-0122122301020032-0331221200123121-3330220113203100-0333000203223013-3010310100131012-3110230230213331-1111020210011003"></a>

<a id="canonical-2312311210310132-0130012230120322-3220201312013012-3212200033203331-3210123030202223-0012230223310331-3121003301121220-3301202003312010"></a>

## proto_redirect property — route_redirect / 200330313210 / 7

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [remove_all_params](resources--http_loadbalancer--reference--group-024.md#canonical-1000301113020132-0332010000322121-3312012101100332-3010130312230213-1112330312023330-3112111232332130-1211323030323323-1123232123230311): complete subsection reference.

<a id="canonical-1023222221310302-2102203233103330-0323200223311010-0212110220023120-1033230310031202-3112332211111011-1130332033002022-3233210312010331"></a>

<a id="canonical-0001123112222332-3012111101223020-3031132230210222-3212203210001331-0223031330311123-3000132212210310-2011100000332030-1001003110120322"></a>

## replace_params property — route_redirect / 200330313210 / 8

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0323203311130202-0222003013111303-2001231330113223-1333133011232001-3233131332013212-0211120033113321-0332302021000112-2300132011032111"></a>

<a id="canonical-3111312030312023-0302112111221322-1320332001310120-2031021201111010-3303203020130220-0112131111123003-3313031112231213-2121233323110210"></a>

## response_code property — route_redirect / 200330313210 / 9

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [retain_all_params](resources--http_loadbalancer--reference--group-024.md#canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020): complete subsection reference.

<a id="canonical-0002013220011303-3311111033022302-2320132210110030-0333200302220011-0031031033023300-1030221010033012-0232313111011121-2021112203331030"></a>

## Next pages — route_redirect / 200330313210 / 10

- [routes.redirect_route.route_redirect.remove_all_params](resources--http_loadbalancer--reference--group-024.md#canonical-1000301113020132-0332010000322121-3312012101100332-3010130312230213-1112330312023330-3112111232332130-1211323030323323-1123232123230311)
- [routes.redirect_route.route_redirect.retain_all_params](resources--http_loadbalancer--reference--group-024.md#canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1000301113020132-0332010000322121-3312012101100332-3010130312230213-1112330312023330-3112111232332130-1211323030323323-1123232123230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031211001000200-0120003231322313-0001303230032320-3311230112033010-0100223111213221-0210033203103100-3013132000212232-1221101102012221"></a>

## routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 223233321131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-2233231301221232-2231323032022230-2310032121130031-3212033001331233-0212033220332231-2322220323010233-0112321312211321-1001211312131201"></a>

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

<a id="canonical-0122212112100122-3131213021133202-1323033231312130-3100131311201213-2030120020132211-3113122212010030-1100103213131120-1003010022001233"></a>

## Direct properties — remove_all_params / 223233321131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010300133320332-1231230300201123-3333333003023013-2121203312001223-2330220121223310-3112321312331231-0021131313221223-1221223210010213"></a>

## Next pages — remove_all_params / 223233321131 / 4

- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202020103122310-3122301223132022-1112003211330031-3202213310301130-1013323221230222-1223023030301101-2320012133130200-1003130003320012"></a>

## routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 120011202323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-024.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-0030323002110330-2012303220102223-2001303002122122-2131131021311312-1131313200202323-3310323321332222-1310100010323123-3010233222321222"></a>

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

<a id="canonical-1332222100220201-1310210120311202-1211312000010133-3300203310322223-1230111312222102-0130100021003200-0201300222110323-0113223002121230"></a>

## Direct properties — retain_all_params / 120011202323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312021003012322-2200313132000030-1030200023001002-3111220303031212-0202112213032000-0001003301030333-0313333300000223-0210201012023103"></a>

## Next pages — retain_all_params / 120011202323 / 4

- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-024.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320110210012020-1011001131023113-3122012023013100-0202021233102001-1103221323310123-0220222210220203-3200213033201323-2003233102131121"></a>

## routes.route_state_disabled — route_state_disabled / 302111233132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.route_state_disabled

<a id="canonical-3230110113113210-3221321311303201-3311022102121122-0111130013321032-1012101230021333-0002000231013000-2303303123123213-0220133122111332"></a>

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
route_state_disabled = {}
```

<a id="canonical-3211202211203333-2122201213101001-1220301011220033-3131332032002331-3112101120323022-2112323021010013-3220211030120112-2030330330010222"></a>

## Direct properties — route_state_disabled / 302111233132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003330110233233-0101213033000320-2203312311101313-2100130121321230-0023030110322321-2002131222132201-1300011230331222-1013002311131322"></a>

## Next pages — route_state_disabled / 302111233132 / 4

- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012022231302000-2213030003001303-0222232321213210-2301323013110123-2020230220310130-2312102002210321-2300230301232010-0220023302312233"></a>

## routes.route_state_enabled — route_state_enabled / 231131011023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.route_state_enabled

<a id="canonical-1310132031110303-0300023231112102-1321100032211310-2200031220223023-3002331202033120-1010111121212223-1320113121302113-0222203230210311"></a>

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
route_state_enabled = {}
```

<a id="canonical-0211323323313200-0003122220212103-1213032132122302-0030130133001001-0320222023130003-1311120313020322-1332302331311011-2030001101312313"></a>

## Direct properties — route_state_enabled / 231131011023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132010300131113-0222303001110113-2132113102101133-3210310001121301-1313212131232200-0231311312203310-0021231311210312-0001313310010112"></a>

## Next pages — route_state_enabled / 231131011023 / 4

- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033323230021223-3231020102311122-3001112132211232-1020230020113130-2132033201010022-0031001103120001-2223132022032212-2233122031132000"></a>

## routes.simple_route — simple_route / 331000110202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.simple_route

<a id="canonical-3222110311023220-2231131202300221-3113002110020033-0131333113222113-3312223323223331-2020231333222223-0311232301133203-1211320210323332"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Upstream description:

A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_pools"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit"),
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012123000302111-0030300303002112-0311311223211320-1031112300131103-2132223212121220-2021032310103110-3231011001313200-2200211230221320"></a>

## Direct properties — simple_route / 331000110202 / 3

- [advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311): complete subsection reference.

- [auto_host_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-2301322211233333-2220203210213133-1223123100323130-0320033102000201-0232100103031301-3131223110111100-0211212220020322-2122231131101230): complete subsection reference.

- [caching_disable](resources--http_loadbalancer--reference--group-026.md#canonical-0202222022121123-0202330031011312-2121330022012301-3130103232111220-2101020132131122-0302212100030102-3222333202231011-0111032222221202): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-026.md#canonical-3030300201301203-0333111031333222-0020021113110221-1230233101000133-3230333211121131-3112030300111101-3132120113232101-3332020012003023): complete subsection reference.

- [disable_host_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-0302130301011103-0313021230021220-0323232031210001-2220233310233323-1010133000201122-1213112103132032-0112010123003213-3123321131203022): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-026.md#canonical-3031023200230323-1323321113310303-2300020320132130-1030230022312313-0003202303001303-0002123020321212-3203111333003001-1323231012303111): complete subsection reference.

<a id="canonical-2022011301231211-3231123220011312-3110303300310302-1133203231130111-2212333333130313-0220202331130322-3232123330102010-1211203312131202"></a>

<a id="canonical-1330000111233212-0122310320332311-1222330022302200-2200202202023231-3101321101020111-1002033130133122-0310330313100021-0200333011230232"></a>

## host_rewrite property — simple_route / 331000110202 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3001120032033210-0220211231011001-2100222231322013-3302013033103131-0130203113210303-3313032331002022-3302023133310032-1102102003222223"></a>

<a id="canonical-1330321113232012-3211221103301013-0033221323310212-1302121002133120-2010010311332111-1311320231010123-1023020121030330-1230111031132113"></a>

## http_method property — simple_route / 331000110202 / 5

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

- [incoming_port](resources--http_loadbalancer--reference--group-026.md#canonical-2021321033003230-1330132333221000-2002303112001020-2221033213330222-2123322230122213-3213310100232300-3232300311302131-2032221321131323): complete subsection reference.

- [origin_pools](resources--http_loadbalancer--reference--group-026.md#canonical-2130022000022203-0210102312021103-2132200221323000-1322100313301310-0001331313213000-0231331120331111-0230300123302313-2012003132322232): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-026.md#canonical-3313101103010332-3202233212103301-1302311031332233-0323313323113103-1302101301011133-3132022320301012-2320322112030300-0130020323203110): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-026.md#canonical-1321000220322201-0133323231302021-1011033331200313-1120000203111001-2032133211021021-1010223322202133-2132131102032110-2130300320231101): complete subsection reference.

<a id="canonical-1203133303210130-0120000023111001-0222033111112112-3120020213010203-2332210110313312-0332111001011132-1302030031131213-2123311001232003"></a>

## Next pages — simple_route / 331000110202 / 6

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.auto_host_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-2301322211233333-2220203210213133-1223123100323130-0320033102000201-0232100103031301-3131223110111100-0211212220020322-2122231131101230)
- [routes.simple_route.caching_disable](resources--http_loadbalancer--reference--group-026.md#canonical-0202222022121123-0202330031011312-2121330022012301-3130103232111220-2101020132131122-0302212100030102-3222333202231011-0111032222221202)
- [routes.simple_route.caching_inherit](resources--http_loadbalancer--reference--group-026.md#canonical-3030300201301203-0333111031333222-0020021113110221-1230233101000133-3230333211121131-3112030300111101-3132120113232101-3332020012003023)
- [routes.simple_route.disable_host_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-0302130301011103-0313021230021220-0323232031210001-2220233310233323-1010133000201122-1213112103132032-0112010123003213-3123321131203022)
- [routes.simple_route.headers](resources--http_loadbalancer--reference--group-026.md#canonical-3031023200230323-1323321113310303-2300020320132130-1030230022312313-0003202303001303-0002123020321212-3203111333003001-1323231012303111)
- [routes.simple_route.incoming_port](resources--http_loadbalancer--reference--group-026.md#canonical-2021321033003230-1330132333221000-2002303112001020-2221033213330222-2123322230122213-3213310100232300-3232300311302131-2032221321131323)
- [routes.simple_route.origin_pools](resources--http_loadbalancer--reference--group-026.md#canonical-2130022000022203-0210102312021103-2132200221323000-1322100313301310-0001331313213000-0231331120331111-0230300123302313-2012003132322232)
- [routes.simple_route.path](resources--http_loadbalancer--reference--group-026.md#canonical-3313101103010332-3202233212103301-1302311031332233-0323313323113103-1302101301011133-3132022320301012-2320322112030300-0130020323203110)
- [routes.simple_route.query_params](resources--http_loadbalancer--reference--group-026.md#canonical-1321000220322201-0133323231302021-1011033331200313-1120000203111001-2032133211021021-1010223322202133-2132131102032110-2130300320231101)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030002220330232-2030332021332122-0310021123300320-0231033031222313-2002201203222230-0131012230231221-1233031021012301-3110103330133311"></a>

## routes.simple_route.advanced_options — advanced_options / 330021313210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- routes.simple_route.advanced_options

<a id="canonical-3112132230020132-0112021320130032-2322112130310310-2302312100111023-3123231311312103-3010303303021113-1013311031322010-3312032000310133"></a>

Type: `"object"`. single nested block, Optional.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingObjectAttributes("buffer_policy",
    "common_buffering"),
  validators.ConflictingObjectAttributes("common_hash_policy",
    "specific_hash_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "no_retry_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("disable_mirroring",
    "mirror_policy"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "regex_rewrite"),
  validators.ConflictingObjectAttributes("disable_spdy",
    "enable_spdy"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("disable_web_socket_config",
    "web_socket_config"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingObjectAttributes("no_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133213132231302-3111302300301320-2031101231221330-1303213210000322-1200112100230203-2003211210110303-0320223001032232-1331301000002001"></a>

## Direct properties — advanced_options / 330021313210 / 3

- [app_firewall](resources--http_loadbalancer--reference--group-024.md#canonical-2001003120333132-1102300103312302-3101121230003012-2130113322120203-1111311330221133-1333123230101210-0333012200031023-2230100113201111): complete subsection reference.

- [bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121): complete subsection reference.

- [buffer_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1211001130101123-0202021333221102-1013312002223231-0033202013223031-2322012302201213-0221112223001103-3230112102023130-3102232202213203): complete subsection reference.

- [common_buffering](resources--http_loadbalancer--reference--group-025.md#canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032): complete subsection reference.

- [common_hash_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--reference--group-025.md#canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311): complete subsection reference.

- [default_retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311): complete subsection reference.

<a id="canonical-1020001231020323-2032232202330032-1212122323221322-3001000133101120-2320330013320230-1023131330100220-1233110321231330-3231112002200232"></a>

<a id="canonical-2212000212200030-3312330103231101-3212201231202222-1031233013202332-1232303022222123-3003103022012230-3302010331100331-1013100220300212"></a>

## disable_location_add property — advanced_options / 330021313210 / 4

Type: `"bool"`. Optional.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [disable_mirroring](resources--http_loadbalancer--reference--group-025.md#canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232): complete subsection reference.

- [disable_prefix_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123): complete subsection reference.

- [disable_spdy](resources--http_loadbalancer--reference--group-025.md#canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--reference--group-025.md#canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111): complete subsection reference.

- [disable_web_socket_config](resources--http_loadbalancer--reference--group-025.md#canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030): complete subsection reference.

- [do_not_retract_cluster](resources--http_loadbalancer--reference--group-025.md#canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333): complete subsection reference.

- [enable_spdy](resources--http_loadbalancer--reference--group-025.md#canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-025.md#canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131): complete subsection reference.

- [inherited_bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-025.md#canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000): complete subsection reference.

- [inherited_waf](resources--http_loadbalancer--reference--group-025.md#canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323): complete subsection reference.

- [inherited_waf_exclusion](resources--http_loadbalancer--reference--group-025.md#canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231): complete subsection reference.

- [mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023): complete subsection reference.

- [no_retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111): complete subsection reference.

<a id="canonical-0202322133003233-2323130221211011-0331013321202102-3001310231320022-3020301123122111-2201333333123123-0011330321020221-2001302303112231"></a>

<a id="canonical-0011200312121123-2033031121131330-1013210112112130-0033131022120212-1310131330231013-2303101112003311-0323030213011000-0233323203033322"></a>

## prefix_rewrite property — advanced_options / 330021313210 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_prefix\_rewrite regular expression\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regular expression path
matching, the entire path (not including the query string) will be swapped with this value.

Upstream description:

Exclusive with \[disable\_prefix\_rewrite regular expression\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regular expression path
matching, the entire path (not including the query string) will be swapped with this value.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0110321121101133-0021201220321312-0111110232023013-1112120320000010-1310323201022303-1002021330102010-3022330211023330-0303323131312031"></a>

<a id="canonical-1233001313000133-3230101231220232-2020011333233132-2311102000311000-1120211020113113-3023313130330000-1130113000110131-2232202331210233"></a>

## priority property — advanced_options / 330021313210 / 6

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [regex_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321): complete subsection reference.

<a id="canonical-3003321220100110-0303003032212122-1111131323131123-2301133021010232-3233102102201030-1113103310301201-2211113220032032-1320320233220122"></a>

<a id="canonical-0321130310323300-3310232201231112-1131202000313112-0121312030213200-1123100031301033-1032303200320230-3012103322123030-1201011203102331"></a>

## request_cookies_to_remove property — advanced_options / 330021313210 / 7

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313): complete subsection reference.

<a id="canonical-2103233003313121-3222000123220312-2101323120100300-2320100202000003-0201221132200213-2002032213002230-0321321121021002-1230130213102030"></a>

<a id="canonical-1100122123130212-2302021030110233-1303022231311232-1133331303000200-1121011113332331-3230321023021012-3102110130021013-0003311202333212"></a>

## request_headers_to_remove property — advanced_options / 330021313210 / 8

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020): complete subsection reference.

<a id="canonical-3131112013031110-1121310200030031-3331302220312312-3001332110312231-2320310001313021-1311102032312211-2013102033323303-0132320333132331"></a>

<a id="canonical-2233121130121333-2122333232221031-2311222303113203-3221302102102112-1013330103221323-1200320301303123-2022221203313030-1221032020301111"></a>

## response_cookies_to_remove property — advanced_options / 330021313210 / 9

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230): complete subsection reference.

<a id="canonical-0213021013133312-0203013313301033-0331312100000003-3103111103102000-3130030023222222-1223031020110301-2320330112323122-0302110112300313"></a>

<a id="canonical-2230320203312200-3211103212002021-1231331023003002-0020013132010022-1011131230111212-1230212213200322-1001122210201121-1011222231123230"></a>

## response_headers_to_remove property — advanced_options / 330021313210 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](resources--http_loadbalancer--reference--group-026.md#canonical-1001202133030320-0013120123332230-2010322030001110-0213023101000331-3211331232101131-1213100001010311-0221112031110323-2311021231132111): complete subsection reference.

- [retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2031213021312332-2113132313130100-3222333013111133-0122010300022110-1323302212201331-0133200321222002-3221100221013001-0011102330211230): complete subsection reference.

- [specific_hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2332103311111231-3001202122300101-0003322101013201-2300220101202301-0330222321301302-3233120102200011-0001011001021100-2330220212112132): complete subsection reference.

<a id="canonical-0132110101311213-3022130021113002-0331332231002321-0022322011301222-1102233001102311-0333023130223021-0331031112021222-1111300122223322"></a>

<a id="canonical-3323210301321100-1022333121201322-0313323222311323-0112021033023223-3011222222100212-1200232323000330-0330032311222112-2222002201302212"></a>

## timeout property — advanced_options / 330021313210 / 11

Type: `"number"`. Optional.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Upstream description:

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-026.md#canonical-3123100221313210-3130201120132022-0032331322102220-1220110221031321-3132211020111030-3012322332111331-0210011333220010-2333010312023201): complete subsection reference.

- [web_socket_config](resources--http_loadbalancer--reference--group-026.md#canonical-2221130310300232-3333311201210030-2220032332203101-0111312233003310-0330020033223201-2312212013003202-2330320321012310-3132220122331020): complete subsection reference.

<a id="canonical-3131200023012310-0121200303333330-0311313123023211-1023211313331130-1133220032330230-3011133331232300-3130303330213233-2123121330332103"></a>

## Next pages — advanced_options / 330021313210 / 12

- [routes.simple_route.advanced_options.app_firewall](resources--http_loadbalancer--reference--group-024.md#canonical-2001003120333132-1102300103312302-3101121230003012-2130113322120203-1111311330221133-1333123230101210-0333012200031023-2230100113201111)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- [routes.simple_route.advanced_options.buffer_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1211001130101123-0202021333221102-1013312002223231-0033202013223031-2322012302201213-0221112223001103-3230112102023130-3102232202213203)
- [routes.simple_route.advanced_options.common_buffering](resources--http_loadbalancer--reference--group-025.md#canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032)
- [routes.simple_route.advanced_options.common_hash_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321)
- [routes.simple_route.advanced_options.cors_policy](resources--http_loadbalancer--reference--group-025.md#canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311)
- [routes.simple_route.advanced_options.default_retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311)
- [routes.simple_route.advanced_options.disable_mirroring](resources--http_loadbalancer--reference--group-025.md#canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232)
- [routes.simple_route.advanced_options.disable_prefix_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123)
- [routes.simple_route.advanced_options.disable_spdy](resources--http_loadbalancer--reference--group-025.md#canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012)
- [routes.simple_route.advanced_options.disable_waf](resources--http_loadbalancer--reference--group-025.md#canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111)
- [routes.simple_route.advanced_options.disable_web_socket_config](resources--http_loadbalancer--reference--group-025.md#canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030)
- [routes.simple_route.advanced_options.do_not_retract_cluster](resources--http_loadbalancer--reference--group-025.md#canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333)
- [routes.simple_route.advanced_options.enable_spdy](resources--http_loadbalancer--reference--group-025.md#canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031)
- [routes.simple_route.advanced_options.endpoint_subsets](resources--http_loadbalancer--reference--group-025.md#canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131)
- [routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-025.md#canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000)
- [routes.simple_route.advanced_options.inherited_waf](resources--http_loadbalancer--reference--group-025.md#canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323)
- [routes.simple_route.advanced_options.inherited_waf_exclusion](resources--http_loadbalancer--reference--group-025.md#canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231)
- [routes.simple_route.advanced_options.mirror_policy](resources--http_loadbalancer--reference--group-025.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023)
- [routes.simple_route.advanced_options.no_retry_policy](resources--http_loadbalancer--reference--group-025.md#canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111)
- [routes.simple_route.advanced_options.regex_rewrite](resources--http_loadbalancer--reference--group-025.md#canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332)
- [routes.simple_route.advanced_options.request_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321)
- [routes.simple_route.advanced_options.request_headers_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--reference--group-025.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020)
- [routes.simple_route.advanced_options.response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230)
- [routes.simple_route.advanced_options.retract_cluster](resources--http_loadbalancer--reference--group-026.md#canonical-1001202133030320-0013120123332230-2010322030001110-0213023101000331-3211331232101131-1213100001010311-0221112031110323-2311021231132111)
- [routes.simple_route.advanced_options.retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2031213021312332-2113132313130100-3222333013111133-0122010300022110-1323302212201331-0133200321222002-3221100221013001-0011102330211230)
- [routes.simple_route.advanced_options.specific_hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2332103311111231-3001202122300101-0003322101013201-2300220101202301-0330222321301302-3233120102200011-0001011001021100-2330220212112132)
- [routes.simple_route.advanced_options.waf_exclusion_policy](resources--http_loadbalancer--reference--group-026.md#canonical-3123100221313210-3130201120132022-0032331322102220-1220110221031321-3132211020111030-3012322332111331-0210011333220010-2333010312023201)
- [routes.simple_route.advanced_options.web_socket_config](resources--http_loadbalancer--reference--group-026.md#canonical-2221130310300232-3333311201210030-2220032332203101-0111312233003310-0330020033223201-2312212013003202-2330320321012310-3132220122331020)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2001003120333132-1102300103312302-3101121230003012-2130113322120203-1111311330221133-1333123230101210-0333012200031023-2230100113201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121010303031121-0103131211112010-2123101011030232-0332003101132021-1111133310200320-3023320133312210-2221231011232101-3000333221202210"></a>

## routes.simple_route.advanced_options.app_firewall — app_firewall / 130100302232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.app_firewall

<a id="canonical-1000212213232202-2100301113231300-2120101013333023-3331233232132321-0311303111131302-1212002120023023-3331301001120301-3012202322222021"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233310011231200-0233003231330023-1011031201123121-0211030003321033-1131133013212200-3320212012131003-1030111322121331-3201023110311031"></a>

## Direct properties — app_firewall / 130100302232 / 3

<a id="canonical-0111132101120012-1223113220311220-0022113222213022-3231103021001013-1011300213303331-3220013012331202-1130222023001310-2320213331113220"></a>

<a id="canonical-3332201333001012-3221103132312132-3100201123212211-0111010132332113-0000021003303323-0323101302201111-0330201333133221-0121333233013120"></a>

## name property — app_firewall / 130100302232 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2121202211030001-0212121210021011-2332322233333113-1112203310200223-3201013232011122-0032131333121110-0020310113320132-0311313103013113"></a>

<a id="canonical-3030112210320031-3001231231102310-2012000111212200-3110302103003010-2310021012213133-2211002103322223-2313231123030000-3230213122123332"></a>

## namespace property — app_firewall / 130100302232 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1210330122100320-1202022131230121-3033220032101300-2023233120323223-1013333221331013-3120133021200033-0203233313222022-0323103113121130"></a>

<a id="canonical-1002332032110310-2232001220011202-3102110313100013-3300330300023233-3033121001112210-0013002000110003-0100221030101000-1332332131120322"></a>

## tenant property — app_firewall / 130100302232 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0010120122010310-3233132211012030-0212312333022001-0002121000323311-1032310201001211-0010100213111010-0320103303302221-2001233303112232"></a>

## Next pages — app_firewall / 130100302232 / 7

- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011121133213003-2003233313002233-1320110321300012-2313213330110300-3130030002120232-0221212333210120-2221200011310302-3221100031313131"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection — bot_defense_javascript_injection / 012203300103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.bot_defense_javascript_injection

<a id="canonical-3112222331022130-1332123020020302-2300331200133303-1200121303103320-3010212320333220-1301323113301132-2333032212133323-1212203233121122"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense JavaScript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210012200001031-2132023010023120-1111132300011020-0303032102020231-0011220122102120-0001011210312001-0100230020030030-2131120223233201"></a>

## Direct properties — bot_defense_javascript_injection / 012203300103 / 3

<a id="canonical-2010121222013120-0130321033322012-1003222230200300-3131022000222030-3202012002031333-1103223332030021-1031311110031301-0031221300113021"></a>

<a id="canonical-2213101020200132-1132003011031012-0002230101033203-2011312020112211-3320021322102320-2333223303001120-0023101123302132-3111310020112220"></a>

## javascript_location property — bot_defense_javascript_injection / 012203300103 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121): complete subsection reference.

<a id="canonical-2131003131332010-2332103310011032-3200301023330222-0031100012221311-1111330011100023-0301321210232213-0322303131201233-0020100101003301"></a>

## Next pages — bot_defense_javascript_injection / 012203300103 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](resources--http_loadbalancer--reference--group-024.md#canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0230323012012233-1331313230201303-2013102231310211-2300022200010131-0013100032031013-1323313021311010-2301021330031011-3003103033102121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222233101211003-1121221212013303-2021021231310322-2030221330320022-2130012111123300-3323202023201300-2333220033200131-0011201112001032"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags — javascript_tags / 303210132201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-024.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-024.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-024.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

<a id="canonical-3211033200233212-2313002301301112-2303030310230002-1210212112123000-2303001121000111-1213321320221211-2030132022310021-2302001311333213"></a>

Type: `"object"`. list nested block, Optional.

Select Add item to configure your JavaScript tag. If adding both Bot Adv and Fraud, the Bot
JavaScript should be added first.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("javascript_url")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
javascript_tags {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031323100032321-3130010301210212-0131221032133003-2212110100100010-2303111222000011-0133333203311020-2102011302033033-0012122212331121"></a>

## Direct properties — javascript_tags / 303210132201 / 3

<a id="canonical-1303311102311322-2132132101111330-1010012220010001-2011233130120010-0300310013022233-3230022020011021-3132203320111113-3112123311103300"></a>

<a id="canonical-1130321122222312-2212013110221031-3113312311332310-3221002333222203-3022211312212321-1121203210022132-3023320333130130-2122312031322322"></a>

## javascript_url property — javascript_tags / 303210132201 / 4

Type: `"string"`. Optional.

Please enter the full URL (include domain and path), or relative path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](resources--http_loadbalancer--reference--group-024.md#canonical-2113101232122210-0112022221131010-2302331003003113-1033122321301121-2100233130331013-1003333011321300-2220033023013303-1121111302131110): complete subsection reference.

<a id="canonical-0001231203332102-0301110202202210-0032122311132003-0132120010211033-0132302002202323-2332301332221221-1322313210310202-0222111223010300"></a>

## Next pages — javascript_tags / 303210132201 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--http_loadbalancer--reference--group-024.md#canonical-2113101232122210-0112022221131010-2302331003003113-1033122321301121-2100233130331013-1003333011321300-2220033023013303-1121111302131110)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-024.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2113101232122210-0112022221131010-2302331003003113-1033122321301121-2100233130331013-1003333011321300-2220033023013303-1121111302131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
