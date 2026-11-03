# Reference
## Account
<details><summary><code>client.Account.Get() -> *affinity.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read the organization and Test or Live mode selected by the API key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
_, err := client.Account.Get(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Feedback
<details><summary><code>client.Feedback.Create(request) -> *affinity.CreateFeedbackResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send feedback without credentials, or with an API key to include its organization and mode. Do not include patient information. API-key callers are limited to 10 submissions per hour per key. Anonymous callers are limited to 10 submissions per hour per IP address. A 429 response includes Retry-After in seconds.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.CreateFeedbackRequest{
    IdempotencyKey: "Idempotency-Key",
    Message: "The error does not explain which field is invalid.",
    PageURL: affinity.String(
        "https://docs.affinityrx.com/errors/",
    ),
    RequestID: affinity.String(
        "6e562761-eff3-4626-a8b9-25ce7c670ccd",
    ),
}
_, err := client.Feedback.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**message:** `string` — Text with at least one visible character.
    
</dd>
</dl>

<dl>
<dd>

**pageURL:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requestID:** `*string` — The `requestId` of the response, as the API returned it: a UUID or its `req_` form.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Medications
<details><summary><code>client.Medications.List() -> *affinity.MedicationList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Every medication you can order in the current mode, A to Z by name, then by ID. Send query to find medications by name: every word must appear in the name, so check the spelling when nothing matches. Then choose by name, strength, form and pharmacy. Send practiceId to list what that practice can order, at its prices; without it, the caller's organization is used. Prices are in cents. `requirements` names the optional prescription fields this pharmacy requires; `compoundingReasons` and `quantities` list the accepted values when the pharmacy limits them, counted in `quantityUnit`. With a `fill`, `price` is per fill and each quantity counts `quantityUnit`, not fills: it is a whole number of fills in units. For a fill of 30 tablets, quantity 60 is two fills. For vials, quantity 1 is one vial and quantity 2 is two vials. A 2 mL vial lists quantities 1–10; its volume is separate from its quantity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListMedicationsRequest{
    PracticeID: affinity.String(
        "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    ),
}
_, err := client.Medications.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**practiceID:** `*string` — Quote prices and availability for this practice. A platform may select only its own child practices; a practice may select only itself. Omit to use the caller's organization.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum records per page, from 1 to 100. Defaults to 25.
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` — Page forward using the previous response's nextCursor. Keep the same filters, sort, organization and mode.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Match medication names that contain every word of this text, in any order, ignoring case and accents. Each word must appear in the name, so a misspelled word matches nothing. Strength, form and pharmacy are not searched. An empty query lists every medication. Restart pagination after changing query.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Medications.Get(MedicationID) -> *affinity.GetMedicationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one available medication, its pricing, and pharmacy details for the current organization and mode. Pass expand=variants to include its other strengths, forms and pharmacies.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.GetMedicationsRequest{
    MedicationID: "med_01j2y8m6jcc9tt24af5pw9x1bc",
    PracticeID: affinity.String(
        "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    ),
}
_, err := client.Medications.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**medicationID:** `string` — The medication ID.
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` — Quote prices and availability for this practice. A platform may select only its own child practices; a practice may select only itself. Omit to use the caller's organization.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*affinity.GetMedicationsRequestExpandItem` — Add related records to the response. `variants` adds the medication's other strengths, forms and pharmacies.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Orders
<details><summary><code>client.Orders.List() -> *affinity.OrderList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List accessible orders in the current mode, newest first. `externalOrderId` matches exactly, including case. `updatedAfter` accepts an ISO 8601 timestamp with a timezone and returns strictly later updates. Keep filters fixed while paging. `sort` takes `patient` (lowercase firstName + space + lastName A to Z, then order ID) or `created` (oldest first); a `_desc` suffix reverses. Creation time or patient name and the order ID determine cursor order, even when updatedAt changes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListOrdersRequest{
    PatientID: affinity.String(
        "pat_01j2y8m6jcc9tt24af5pw9x1bc",
    ),
    PracticeID: affinity.String(
        "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    ),
}
_, err := client.Orders.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Maximum records per page, from 1 to 100. Defaults to 25.
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` — Page forward using the previous response's nextCursor. Keep the same filters, sort, organization and mode.
    
</dd>
</dl>

<dl>
<dd>

**externalOrderID:** `*string` — Match the external order ID exactly, including case. Combine with practiceId when your platform reuses IDs across practices.
    
</dd>
</dl>

<dl>
<dd>

**updatedAfter:** `*time.Time` — Return orders whose updatedAt is strictly later than this ISO 8601 timestamp. Sorting stays by creation time and ID unless sort is supplied. Keep this cutoff unchanged while paging.
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `*string` — Limit results to the patient with this pat_ ID.
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` — Limit results to the practice with this prac_ ID, within your authorized practices.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Search saved order snapshots: patient names, email, phone and external ID; order external ID; prescriber name and NPI; medication names and brands; pharmacy name. Whole ord_ and pat_ IDs and six-character order numbers also match. Text matches ignore case and allow substrings, word prefixes and similar spellings for alphanumeric words of four or more characters. Every one of the first eight words containing letters or numbers must match; punctuation-only words are ignored. An empty or punctuation-only query leaves the list unfiltered. Date of birth, separate strength fields, practice names, prac_ IDs and rx_ IDs are not searched. Restart pagination when filters change.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListOrdersRequestStatus` — Return only orders with this status. Omit to include all statuses.
    
</dd>
</dl>

<dl>
<dd>

**cancelRequested:** `*affinity.ListOrdersRequestCancelRequested` — true returns only submitted or processing orders with a cancel the pharmacy has not confirmed yet (cancelRequestedAt set). false leaves those orders out. Omit to include both.
    
</dd>
</dl>

<dl>
<dd>

**recentlyShipped:** `*affinity.ListOrdersRequestRecentlyShipped` — true returns only orders whose status is shipped and that shipped in the last 60 days; delivered orders are excluded. false leaves those orders out. Omit to include both.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*affinity.ListOrdersRequestSort` — patient sorts the lowercase text firstName + space + lastName A to Z, then order ID; created sorts oldest first. A _desc suffix reverses the order. Defaults to created_desc. Restart pagination after changing sort.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Create(request) -> *affinity.Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

One request creates the order, signs it and sends it to the pharmacy. Platforms must send practiceId; practice callers can omit it to use their own practice. With an API key, name the prescriber and send signatureAttestation. Optional prescriber phone and address each default to the practice; no account is looked up by NPI. In the app, the signed-in member signs with the NPI in their settings, or names the prescriber with the same attestation. With patient.id (a pat_ ID) or patient.externalId, the order reuses that patient of the practice and mode and updates the saved record; without them, each order creates a new patient. If id and externalId identify different patients, the request returns 422 invalid_order. When the pharmacy already holds the patient's name, a different firstName or lastName returns 422 invalid_order. Live orders require Live access on the practice and its platform. No payment method is required.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.OrderInput{
    IdempotencyKey: "Idempotency-Key",
    Patient: &affinity.PatientInput{
        Address: &affinity.AddressInput{
            Line1: "line1",
            City: "city",
            State: "state",
            PostalCode: "postalCode",
        },
        Allergies: []*affinity.PatientInputAllergiesItem{
            &affinity.PatientInputAllergiesItem{
                Category: affinity.PatientInputAllergiesItemCategoryDrug,
                Display: "display",
            },
        },
        DateOfBirth: "dateOfBirth",
        Email: "email",
        FirstName: "firstName",
        Sex: affinity.PatientInputSexFemale,
        LastName: "lastName",
        Phone: "phone",
    },
    Prescriptions: []*affinity.PrescriptionInput{
        &affinity.PrescriptionInput{
            Directions: "directions",
            MedicationID: "med_01j2y8m6jcc9tt24af5pw9x1bc",
            Quantity: 1,
            Refills: 1,
        },
    },
    Shipping: &affinity.OrderInputShipping{
        OptionID: "shp_01j2y8m6jcc9tt24af5pw9x1bc",
    },
    SignatureAttestation: true,
}
_, err := client.Orders.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**externalOrderID:** `*string` — Your order identifier, unique within the practice and mode. A failed or cancelled order keeps it, so send a new one for a replacement order.
    
</dd>
</dl>

<dl>
<dd>

**patient:** `*affinity.PatientInput` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` — Required with a platform API key. Send a child practice ID. Independent practice keys can omit this field to use their own practice.
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.PrescriberInput` 
    
</dd>
</dl>

<dl>
<dd>

**prescriptions:** `[]*affinity.PrescriptionInput` — The 1 to 20 prescriptions approved for this patient.
    
</dd>
</dl>

<dl>
<dd>

**shipping:** `*affinity.OrderInputShipping` — The shipping option and destination.
    
</dd>
</dl>

<dl>
<dd>

**signatureAttestation:** `bool` — Must be true to attest that the named prescriber approved and signed these prescriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Get(OrderID) -> *affinity.Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one order and its prescriptions in the API key's organization and mode.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.GetOrdersRequest{
    OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
}
_, err := client.Orders.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The order ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.Cancel(OrderID) -> *affinity.Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

200 with status cancelled when the cancel is confirmed. 202 with cancelRequestedAt when the pharmacy has not confirmed yet: you get order.cancelled when it confirms, or, when it refuses, cancelRequestedAt becomes null and you get order.updated. If an earlier ask may have reached the pharmacy, a refusal keeps cancelRequestedAt and sends no event until the pharmacy ends the order (order.cancelled or order.shipped). 409 cannot_cancel when the order is shipped, delivered or failed. 409 cancel_refused when the pharmacy refused the cancel, or when Affinity could not reach the pharmacy; only the second says to retry with a new Idempotency-Key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.CancelOrdersRequest{
    OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
    IdempotencyKey: "Idempotency-Key",
}
_, err := client.Orders.Cancel(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**orderID:** `string` — The order ID.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Practices
<details><summary><code>client.Practices.List() -> *affinity.PracticeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List the platform's practices, newest first. Available to platforms only. `sort` takes `name` or `location` (A to Z, practices without a location last), `mode` (without Live access, then with it) or `created` (oldest first); a `_desc` suffix reverses.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListPracticesRequest{}
_, err := client.Practices.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Maximum records per page, from 1 to 100. Defaults to 25.
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` — Page forward using the previous response's nextCursor. Keep the same filters, sort, organization and mode.
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` — Match the external practice ID exactly, including case. Keep this filter unchanged while paging.
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` — Match practice names containing this text, ignoring case. Restart pagination after changing query.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*affinity.ListPracticesRequestSort` — name and location sort A to Z; mode sorts practices without Live access before those with it; created sorts oldest first. A _desc suffix reverses the order. Defaults to created_desc. Missing locations sort last. Restart pagination after changing sort.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.Create(request) -> *affinity.Practice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Platforms only. Name, phone and address are required. The practice starts in Test; `liveEnabled: true` also allows new Live orders, which needs a platform with Live access and a Live caller: a Live key, or a session in Live mode. Turning Live off later preserves access to existing records. A retry with the same key and body returns the stored first response, even if the practice has changed. Use GET /v1/practices/{practiceId} to read the current practice.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.CreatePracticesRequest{
    IdempotencyKey: "Idempotency-Key",
    Address: &affinity.PracticeAddressInput{
        Line1: "line1",
        City: "city",
        State: "state",
        PostalCode: "postalCode",
    },
    Name: "name",
    Phone: "phone",
}
_, err := client.Practices.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` — Your practice identifier, unique among your platform's active practices.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.PracticeAddressInput` 
    
</dd>
</dl>

<dl>
<dd>

**liveEnabled:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The practice name. PerfectRx prints it: use Windows-1252 characters, without < or >.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `string` — A US phone number: 10 digits, with or without +1 or 1 before them. Spaces, dashes, dots and parentheses are allowed, such as (512) 555-0100. Affinity stores and returns +15125550100.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.Get(PracticeID) -> *affinity.Practice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Read one practice accessible to the API key. A platform can read its own practices.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.GetPracticesRequest{
    PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
}
_, err := client.Practices.Get(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**practiceID:** `string` — The practice ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.Update(PracticeID, request) -> *affinity.Practice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

`liveEnabled` allows or refuses new Live orders. Existing records, cancellation, webhooks and billing keep working. Only its platform changes it; enabling requires platform Live access, and either change needs a Live caller: a Live key (403 live_key_required otherwise) or a session in Live mode (403 live_mode_required otherwise). Live orders send `name`, `phone` and `address` to the pharmacy, so a practice with Live access changes them only in Live mode, with a Live key or a session in Live mode (403 live_mode_required otherwise). A retry with the same key and body returns the stored first response, even if the practice has changed. Use GET /v1/practices/{practiceId} to read the current practice.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.UpdatePracticesRequest{
    PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    IdempotencyKey: "Idempotency-Key",
}
_, err := client.Practices.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**practiceID:** `string` — The practice ID.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` — Your practice identifier, unique among your platform's active practices. Send null to clear it.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.UpdatePracticesRequestAddress` — Affinity merges the address parts you send into the saved address. An omitted part keeps its saved value. To move to an address with no suite, send `"line2": null` with the new `line1`, `city`, `state` and `postalCode`. PerfectRx prints it: use Windows-1252 characters, without < or >. Line 1 and line 2 together: 70 characters or fewer.
    
</dd>
</dl>

<dl>
<dd>

**liveEnabled:** `*bool` — Omit it or send null to leave Live unchanged.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The practice name. PerfectRx prints it: use Windows-1252 characters, without < or >.
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` — A US phone number: 10 digits, with or without +1 or 1 before them. Spaces, dashes, dots and parentheses are allowed, such as (512) 555-0100. Affinity stores and returns +15125550100. Empty text counts as null.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.List() -> *affinity.WebhookEndpointList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List webhook endpoints owned by the API key's organization in its current mode. Signing secrets are not returned.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListWebhooksRequest{}
_, err := client.Webhooks.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**limit:** `*int` — Maximum records per page, from 1 to 100. Defaults to 25.
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` — Page forward using the previous response's nextCursor. Keep the same filters, sort, organization and mode.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Create(request) -> *affinity.CreateWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Platforms and independent practices can register up to 16 public HTTPS endpoints per mode. An API key acts only for its own organization and mode. Every event goes to every endpoint. The signing secret is returned by creation and its idempotent replay for 24 hours.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.CreateWebhooksRequest{
    IdempotencyKey: "Idempotency-Key",
    URL: "url",
}
_, err := client.Webhooks.Create(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` — The public HTTPS URL that receives webhook events.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Remove(WebhookEndpointID) -> *affinity.RemoveWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Events not delivered to it yet are dropped.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.RemoveWebhooksRequest{
    WebhookEndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
}
_, err := client.Webhooks.Remove(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookEndpointID:** `string` — The webhook endpoint ID.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.Update(WebhookEndpointID, request) -> *affinity.WebhookEndpoint</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

The signing secret stays the same.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.UpdateWebhooksRequest{
    WebhookEndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
    IdempotencyKey: "Idempotency-Key",
    URL: "url",
}
_, err := client.Webhooks.Update(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookEndpointID:** `string` — The webhook endpoint ID.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>

<dl>
<dd>

**url:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ResendFailed(WebhookEndpointID) -> *affinity.ResendFailedWebhooksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Queue failed order event deliveries from the last 30 days for this endpoint. Only the current organization and mode apply. Each failed delivery gets up to 12 new attempts. Also make pending deliveries waiting on a later retry due now, without resetting their attempt counts. Keep the same event IDs, payloads, and per-order delivery sequence. The queued count includes both kinds. Pending deliveries already due, delivered events, and one-time test events stay unchanged. Send a unique Idempotency-Key for each request. Reuse that key after an uncertain response to return the original count. A request with a new key returns 0 when no eligible failures or pending deliveries waiting on a later retry remain.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ResendFailedWebhooksRequest{
    WebhookEndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
    IdempotencyKey: "Idempotency-Key",
}
_, err := client.Webhooks.ResendFailed(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**webhookEndpointID:** `string` — The webhook endpoint ID.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` — A unique key per operation, 1 to 255 printable ASCII characters without spaces. Retry the same key and body to avoid repeating the action. The replay response depends on the operation. See https://docs.affinityrx.com/guides/#retry-safely.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

