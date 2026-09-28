# Reference
## Locations
<details><summary><code>client.Locations.ListPracticeLocations(PracticeID) -> *affinity.ListPracticeLocationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires locations:read on a practice key or an authorized platform key. Lists active and archived locations by name, with cursor pagination. Use status to filter. Location records are shared between Test and Live for the same practice.
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
request := &affinity.ListPracticeLocationsRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        StartingAfter: affinity.String(
            "loc_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "loc_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Locations.ListPracticeLocations(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPracticeLocationsRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.CreatePracticeLocation(PracticeID, request) -> *affinity.CreatePracticeLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires locations:write and Idempotency-Key for API keys. Creates an active location with a unique name in this practice. Locations are shared between Test and Live. Use the returned ID for Team location access.
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
request := &affinity.CreatePracticeLocationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Name: "name",
    }
client.Locations.CreatePracticeLocation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**city:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**line1:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**line2:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**postalCode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Optional IANA timezone override. Omit to leave unchanged; null clears it. No timezone is inferred when creating a record.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.GetPracticeLocation(PracticeID, LocationID) -> *affinity.GetPracticeLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires locations:read. Returns one active or archived location in the authorized practice.
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
request := &affinity.GetPracticeLocationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        LocationID: "loc_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Locations.GetPracticeLocation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.UpdatePracticeLocation(PracticeID, LocationID, request) -> *affinity.UpdatePracticeLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires locations:write and Idempotency-Key for API keys. Updates only supplied fields; null clears optional contact and address fields. Archived locations cannot be updated. Changes apply to both Test and Live.
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
request := &affinity.UpdatePracticeLocationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        LocationID: "loc_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Locations.UpdatePracticeLocation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**city:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**country:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**line1:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**line2:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**postalCode:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Optional IANA timezone override. Omit to leave unchanged; null clears it. No timezone is inferred when creating a record.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Locations.ArchivePracticeLocation(PracticeID, LocationID) -> *affinity.ArchivePracticeLocationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires locations:write and Idempotency-Key for API keys. Retains the location and historical associations. Archived locations cannot receive new Team assignments. Repeating archive returns the archived location. Changes apply to both Test and Live.
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
request := &affinity.ArchivePracticeLocationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        LocationID: "loc_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Locations.ArchivePracticeLocation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## API Keys
<details><summary><code>client.APIKeys.CreatePlatformPracticeAPIKey(PracticeID, request) -> *affinity.CreatePlatformPracticeAPIKeyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a practice API key for a connected practice. Requires a platform key with service_keys:write and every requested scope. The practice key uses the platform key's Test or Live mode and cannot outlive it. Requires Idempotency-Key for safe retries; the secret is returned in the encrypted replay response for 24 hours.
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
request := &affinity.CreatePlatformPracticeAPIKeyRequest{
        PracticeID: "practiceId",
        IdempotencyKey: "Idempotency-Key",
        Name: "name",
    }
client.APIKeys.CreatePlatformPracticeAPIKey(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**allowedIps:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*affinity.CreatePlatformPracticeAPIKeyRequestScopesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.APIKeys.GetAPIAccess() -> *affinity.GetAPIAccessResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the subject, mode, and scopes for the API key.
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
client.APIKeys.GetAPIAccess(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Account
<details><summary><code>client.Account.GetAccount() -> *affinity.GetAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the platform organization, request livemode, and effective access. API keys report scopes and the service_key role; dashboard sessions report membership permissions. operatingMode describes organization Live access, not the credential's Test/Live mode.
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
request := &affinity.GetAccountRequest{
        OrgID: affinity.String(
            "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Account.GetAccount(
        context.TODO(),
        request,
    )
}
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

**orgID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Catalog
<details><summary><code>client.Catalog.ListCatalogItems() -> *affinity.ListCatalogItemsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists catalog items for the authenticated account and mode. Use view=medications for priced prescription groups with offer counts, pharmacy counts, and strengths; the default view=offers returns individual offers. Use relatedToCatalogItemId to find offers for the same medication and route. When practiceId is supplied, a practice price overrides the platform price and missing overrides inherit the platform price.
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
request := &affinity.ListCatalogItemsRequest{
        RelatedToCatalogItemID: affinity.String(
            "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        CatalogItemID: affinity.String(
            "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        PharmacyIDs: &affinity.ListCatalogItemsRequestPharmacyIDs{
            String: "pharm_01j2y8m6jcc9tt24af5pw9x1bc",
        },
        EndingBefore: affinity.String(
            "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        OrgID: affinity.String(
            "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        PracticeID: affinity.String(
            "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Catalog.ListCatalogItems(
        context.TODO(),
        request,
    )
}
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

**view:** `*affinity.ListCatalogItemsRequestView` 
    
</dd>
</dl>

<dl>
<dd>

**relatedToCatalogItemID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**catalogKind:** `*affinity.ListCatalogItemsRequestCatalogKind` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*affinity.ListCatalogItemsRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**catalogItemID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**availability:** `*affinity.ListCatalogItemsRequestAvailability` 
    
</dd>
</dl>

<dl>
<dd>

**pharmacyIDs:** `*affinity.ListCatalogItemsRequestPharmacyIDs` 
    
</dd>
</dl>

<dl>
<dd>

**dosageForms:** `*affinity.ListCatalogItemsRequestDosageForms` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**hideControlledSubstances:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**hideUnpriced:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**requirement:** `*affinity.ListCatalogItemsRequestRequirement` 
    
</dd>
</dl>

<dl>
<dd>

**routes:** `*affinity.ListCatalogItemsRequestRoutes` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.ListPharmacies() -> *affinity.ListPharmaciesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists pharmacies available to the authenticated account, including approved invite-only relationships.
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
request := &affinity.ListPharmaciesRequest{
        EndingBefore: affinity.String(
            "pharm_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        OrgID: affinity.String(
            "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        PharmacyID: affinity.String(
            "pharm_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "pharm_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Catalog.ListPharmacies(
        context.TODO(),
        request,
    )
}
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

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orgID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**pharmacyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**shipsToState:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.ListShippingOptions(CatalogItemID) -> affinity.ListShippingOptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an array of at most 50 reviewed shipping services eligible for a catalog item, destination, and API mode. destinationState must be a USPS state or territory code. Each option has one temperature; pharmacy catalog summaries list all supported temperatures.
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
request := &affinity.ListShippingOptionsRequest{
        CatalogItemID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        DestinationState: "destinationState",
    }
client.Catalog.ListShippingOptions(
        context.TODO(),
        request,
    )
}
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

**catalogItemID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**destinationState:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**destinationType:** `*affinity.ListShippingOptionsRequestDestinationType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Catalog.RetrievePrescribingOptions(CatalogItemID) -> *affinity.RetrievePrescribingOptionsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires catalog:read. Returns reviewed SIG presets, guided patterns, quantity constraints and product requirements for a practice and mode. Revisions identify changed defaults. No patient-specific rationale or diagnosis is inferred.
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
request := &affinity.RetrievePrescribingOptionsRequest{
        CatalogItemID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Catalog.RetrievePrescribingOptions(
        context.TODO(),
        request,
    )
}
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

**catalogItemID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Orders
<details><summary><code>client.Orders.ListOrders() -> *affinity.ListOrdersResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListOrdersRequest{
        EndingBefore: affinity.String(
            "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        OrderID: affinity.String(
            "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        PatientID: affinity.String(
            "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        PracticeID: affinity.String(
            "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Orders.ListOrders(
        context.TODO(),
        request,
    )
}
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

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**orderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patientExternalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*affinity.ListOrdersRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListOrdersRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CreateOrder(request) -> *affinity.CreateOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates one unsigned order with 1–20 prescriptions for one patient in one practice. Supply patientId or patient; inline patient creation requires patients:write. Prescriber is optional: select by npi, provider id, or integration-scoped externalId, or leave the draft unassigned until signing. First-use prescriber registration requires team:write. Legacy userId is supported but cannot be combined with prescriber. Idempotency-Key is required.
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
request := &affinity.CreateOrderRequest{
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Prescriptions: []*affinity.CreateOrderRequestPrescriptionsItem{
            &affinity.CreateOrderRequestPrescriptionsItem{
                DaysSupply: 1,
                Dispensing: &affinity.CreateOrderRequestPrescriptionsItemDispensing{},
                Directions: "directions",
                MedicationID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
                Quantity: &affinity.CreateOrderRequestPrescriptionsItemQuantity{
                    CreateOrderRequestPrescriptionsItemQuantityOne: affinity.CreateOrderRequestPrescriptionsItemQuantityOneInfinity,
                },
                QuantityUnit: "quantityUnit",
                Refills: 1,
            },
        },
    }
client.Orders.CreateOrder(
        context.TODO(),
        request,
    )
}
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

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.CreateOrderRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**otcItems:** `[]*affinity.CreateOrderRequestOtcItemsItem` 
    
</dd>
</dl>

<dl>
<dd>

**externalOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*affinity.CreateOrderRequestMetadataValue` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patient:** `*affinity.CreateOrderRequestPatient` 
    
</dd>
</dl>

<dl>
<dd>

**shippingAddressID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriptions:** `[]*affinity.CreateOrderRequestPrescriptionsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetOrder(OrderID) -> *affinity.GetOrderResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.GetOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Orders.GetOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CancelOrder(OrderID, request) -> *affinity.CancelOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requests cancellation. HTTP 200 means the request was handled; check cancellation.status for confirmed, pending, partial, or failed. Only confirmed means the entire order is cancelled. Shipment possession makes a fulfillment cancellation too late.
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
request := &affinity.CancelOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Reason: "reason",
    }
client.Orders.CancelOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.ActOnOrderException(OrderID, ExceptionID, request) -> *affinity.ActOnOrderExceptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Acknowledge, retry, contact, or resolve an order exception in the credential's Test/Live mode. assign_to_me requires a signed-in dashboard user; API keys receive 400 and may use acknowledge instead. Actor headers do not create a dashboard assignee.
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
request := &affinity.ActOnOrderExceptionRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        ExceptionID: "fex_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Action: affinity.ActOnOrderExceptionRequestActionAcknowledge,
    }
client.Orders.ActOnOrderException(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**exceptionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**action:** `*affinity.ActOnOrderExceptionRequestAction` 
    
</dd>
</dl>

<dl>
<dd>

**note:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.ListOrderEvents(OrderID) -> *affinity.ListOrderEventsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListOrderEventsRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        EndingBefore: affinity.String(
            "evt_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "evt_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Orders.ListOrderEvents(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetOrderTestSimulation(OrderID) -> *affinity.GetOrderTestSimulationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:write. Available only in Test mode.
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
request := &affinity.GetOrderTestSimulationRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Orders.GetOrderTestSimulation(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.UpdateOrderTestSimulation(OrderID, request) -> *affinity.UpdateOrderTestSimulationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:write and Idempotency-Key. Configure before submission or queue a valid pharmacy event in manual mode. Events use normal order history and Test webhooks. Live requests are rejected.
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
request := &affinity.UpdateOrderTestSimulationRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Mode: affinity.UpdateOrderTestSimulationRequestModeAutomatic,
        Scenario: affinity.UpdateOrderTestSimulationRequestScenarioSuccessful,
    }
client.Orders.UpdateOrderTestSimulation(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*affinity.UpdateOrderTestSimulationRequestMode` 
    
</dd>
</dl>

<dl>
<dd>

**scenario:** `*affinity.UpdateOrderTestSimulationRequestScenario` 
    
</dd>
</dl>

<dl>
<dd>

**action:** `*affinity.UpdateOrderTestSimulationRequestAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.PreviewOrder(request) -> *affinity.PreviewOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:write and catalog:read. Supply exactly one of patientId, patientExternalId, or inline patient details. External-ID lookup additionally requires patients:read; inline details require patients:write. Resolves defaults and explicit edits for 1–20 prescriptions. Reuses stored patient details when identifiers match; otherwise previews inline details without creating a patient. Complete previews contain an orders.create input. Does not create records, reserve prices, sign, charge or transmit. No idempotency key is required. Creation and signing recheck current requirements.
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
request := &affinity.PreviewOrderRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Prescriptions: []*affinity.PreviewOrderRequestPrescriptionsItem{
            &affinity.PreviewOrderRequestPrescriptionsItem{
                MedicationID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
            },
        },
    }
client.Orders.PreviewOrder(
        context.TODO(),
        request,
    )
}
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

**otcItems:** `[]*affinity.PreviewOrderRequestOtcItemsItem` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patientExternalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**patient:** `*affinity.PreviewOrderRequestPatient` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.PreviewOrderRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**shippingAddressID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalOrderID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriptions:** `[]*affinity.PreviewOrderRequestPrescriptionsItem` 
    
</dd>
</dl>

<dl>
<dd>

**shipping:** `*affinity.PreviewOrderRequestShipping` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.SignOrder(OrderID, request) -> *affinity.SignOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:sign, Idempotency-Key, signatureAttestation, and expectedRevision from the reviewed order. Existing integrations may send expectedVersions instead; supply exactly one. A stale revision returns 409 and requires renewed clinician review. Select prescriber by npi, provider id, or integration-scoped externalId, or inherit the draft's prescriber. First-use registration requires team:write. Actor headers are optional audit metadata with prescriber; legacy userId requires matching clinician actor headers. Signing does not submit to a pharmacy.
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
request := &affinity.SignOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        SignatureAttestation: true,
    }
client.Orders.SignOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.SignOrderRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**signatureAttestation:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*string` — Opaque revision of the complete order prescription set. Send the revision you reviewed as expectedRevision; never replace it automatically after a conflict.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersions:** `[]*affinity.SignOrderRequestExpectedVersionsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.SignAndSubmitOrder(OrderID, request) -> *affinity.SignAndSubmitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:sign, Idempotency-Key, signatureAttestation, and expectedRevision from the reviewed order. Existing integrations may send expectedVersions instead; supply exactly one. A stale revision returns 409 and requires renewed clinician review. Select prescriber by npi, provider id, or externalId, or inherit the draft's prescriber. First-use registration requires team:write. Actor headers are optional with prescriber; legacy userId requires matching clinician actor headers. Signs the complete order, then attempts each submission. Signing remains committed if submission fails. Replay the same key after an uncertain response; retry reported submission failures through Submit order with a new key. Submitted means queued, not pharmacy acceptance.
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
request := &affinity.SignAndSubmitOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        SignatureAttestation: true,
    }
client.Orders.SignAndSubmitOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.SignAndSubmitOrderRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**signatureAttestation:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*string` — Opaque revision of the complete order prescription set. Send the revision you reviewed as expectedRevision; never replace it automatically after a conflict.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersions:** `[]*affinity.SignAndSubmitOrderRequestExpectedVersionsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.SubmitOrder(OrderID, request) -> *affinity.SubmitOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:sign and Idempotency-Key. Queues signed prescriptions after rechecking authorization, signature integrity, billing, and fulfillment eligibility. Track pharmacy acceptance through order reads and webhooks. After a partial failure, retry submission with a new idempotency key; already queued prescriptions are not duplicated.
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
request := &affinity.SubmitOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Orders.SubmitOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.SubmitOrderRequestPrescriber` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.RejectOrder(OrderID, request) -> *affinity.RejectOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:sign and Idempotency-Key. Select a prescriber or inherit the draft's prescriber. Legacy userId requires matching clinician actor headers. Supply expectedRevision from the reviewed order, or expectedVersions for existing integrations. Permanently rejects the complete unsigned order after checking its revision.
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
request := &affinity.RejectOrderRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Reason: "reason",
    }
client.Orders.RejectOrder(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.RejectOrderRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**reason:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*string` — Opaque revision of the complete order prescription set. Send the revision you reviewed as expectedRevision; never replace it automatically after a conflict.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersions:** `[]*affinity.RejectOrderRequestExpectedVersionsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.AddOrderPrescription(OrderID, request) -> *affinity.AddOrderPrescriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:write, Idempotency-Key and expectedRevision from the order being edited. Existing integrations may send expectedVersions instead; supply exactly one. Adds a complete prescription to an unsigned Order and returns all new versions. Omitted actor context defaults to the authenticated service account as a system actor. Patient and prescriber attribution stay fixed. Signed orders cannot be amended through this endpoint. Signing and submission require orders:sign through their separate endpoints.
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
request := &affinity.AddOrderPrescriptionRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Prescription: &affinity.AddOrderPrescriptionRequestPrescription{
            DaysSupply: 1,
            Dispensing: &affinity.AddOrderPrescriptionRequestPrescriptionDispensing{},
            Directions: "directions",
            MedicationID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
            Quantity: &affinity.AddOrderPrescriptionRequestPrescriptionQuantity{
                AddOrderPrescriptionRequestPrescriptionQuantityOne: affinity.AddOrderPrescriptionRequestPrescriptionQuantityOneInfinity,
            },
            QuantityUnit: "quantityUnit",
            Refills: 1,
        },
    }
client.Orders.AddOrderPrescription(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*affinity.AddOrderPrescriptionRequestMetadataValue` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*string` — Opaque revision of the complete order prescription set. Send the revision you reviewed as expectedRevision; never replace it automatically after a conflict.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersions:** `[]*affinity.AddOrderPrescriptionRequestExpectedVersionsItem` 
    
</dd>
</dl>

<dl>
<dd>

**prescription:** `*affinity.AddOrderPrescriptionRequestPrescription` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.UpdateOrderPrescription(OrderID, PrescriptionID, request) -> *affinity.UpdateOrderPrescriptionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires orders:write, Idempotency-Key and expectedRevision from the order being edited. Existing integrations may send expectedVersions instead; supply exactly one. Replaces one prescription with complete medication instructions and returns all new versions. Omitted actor context defaults to the authenticated service account as a system actor. Patient and prescriber attribution stay fixed. Signed orders cannot be amended through this endpoint. Signing and submission require orders:sign through their separate endpoints.
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
request := &affinity.UpdateOrderPrescriptionRequest{
        OrderID: "ord_01j2y8m6jcc9tt24af5pw9x1bc",
        PrescriptionID: "rx_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Prescription: &affinity.UpdateOrderPrescriptionRequestPrescription{
            DaysSupply: 1,
            Dispensing: &affinity.UpdateOrderPrescriptionRequestPrescriptionDispensing{},
            Directions: "directions",
            MedicationID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
            Quantity: &affinity.UpdateOrderPrescriptionRequestPrescriptionQuantity{
                UpdateOrderPrescriptionRequestPrescriptionQuantityOne: affinity.UpdateOrderPrescriptionRequestPrescriptionQuantityOneInfinity,
            },
            QuantityUnit: "quantityUnit",
            Refills: 1,
        },
    }
client.Orders.UpdateOrderPrescription(
        context.TODO(),
        request,
    )
}
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

**orderID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriptionID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]*affinity.UpdateOrderPrescriptionRequestMetadataValue` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expectedRevision:** `*string` — Opaque revision of the complete order prescription set. Send the revision you reviewed as expectedRevision; never replace it automatically after a conflict.
    
</dd>
</dl>

<dl>
<dd>

**expectedVersions:** `[]*affinity.UpdateOrderPrescriptionRequestExpectedVersionsItem` 
    
</dd>
</dl>

<dl>
<dd>

**prescription:** `*affinity.UpdateOrderPrescriptionRequestPrescription` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Orders.CreateOrderBatch(request) -> *affinity.CreateOrderBatchResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates 1–20 orders for distinct patients in one practice, each with 1–20 prescriptions. Each accepts patientId or inline patient details. Orders and newly created patients commit atomically; any failure saves none. Requires orders:write and Idempotency-Key; inline patients also require patients:write. Omitted actor context defaults to the authenticated service account as a system actor. Sign and submit each resulting order separately using orders:sign.
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
request := &affinity.CreateOrderBatchRequest{
        IdempotencyKey: "Idempotency-Key",
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        Orders: []*affinity.CreateOrderBatchRequestOrdersItem{
            &affinity.CreateOrderBatchRequestOrdersItem{
                Prescriptions: []*affinity.CreateOrderBatchRequestOrdersItemPrescriptionsItem{
                    &affinity.CreateOrderBatchRequestOrdersItemPrescriptionsItem{
                        DaysSupply: 1,
                        Dispensing: &affinity.CreateOrderBatchRequestOrdersItemPrescriptionsItemDispensing{},
                        Directions: "directions",
                        MedicationID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
                        Quantity: &affinity.CreateOrderBatchRequestOrdersItemPrescriptionsItemQuantity{
                            CreateOrderBatchRequestOrdersItemPrescriptionsItemQuantityOne: affinity.CreateOrderBatchRequestOrdersItemPrescriptionsItemQuantityOneInfinity,
                        },
                        QuantityUnit: "quantityUnit",
                        Refills: 1,
                    },
                },
            },
        },
    }
client.Orders.CreateOrderBatch(
        context.TODO(),
        request,
    )
}
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

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriber:** `*affinity.CreateOrderBatchRequestPrescriber` 
    
</dd>
</dl>

<dl>
<dd>

**orders:** `[]*affinity.CreateOrderBatchRequestOrdersItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Webhooks
<details><summary><code>client.Webhooks.ListWebhookEndpoints() -> *affinity.ListWebhookEndpointsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:read. Returns endpoints owned by the key organization, or the organization selected with X-Affinity-Organization-Id. Platform delegation requires a webhook grant in the key's mode.
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
request := &affinity.ListWebhookEndpointsRequest{
        EndingBefore: affinity.String(
            "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Webhooks.ListWebhookEndpoints(
        context.TODO(),
        request,
    )
}
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

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.CreateWebhookEndpoint(request) -> *affinity.CreateWebhookEndpointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:write and Idempotency-Key. Defaults to the API key organization. A platform can select a practice or pharmacy owner with X-Affinity-Organization-Id and an explicit webhook grant. For platform-owned endpoints, practiceIds narrows delivery to selected connected practices. An empty filter receives all otherwise-authorized events.
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
request := &affinity.CreateWebhookEndpointRequest{
        IdempotencyKey: "Idempotency-Key",
        URL: "url",
    }
client.Webhooks.CreateWebhookEndpoint(
        context.TODO(),
        request,
    )
}
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

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**payloadStyle:** `*affinity.CreateWebhookEndpointRequestPayloadStyle` 
    
</dd>
</dl>

<dl>
<dd>

**subscribedEvents:** `[]*affinity.CreateWebhookEndpointRequestSubscribedEventsItem` 
    
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

<details><summary><code>client.Webhooks.DeleteWebhookEndpoint(EndpointID) -> *affinity.DeleteWebhookEndpointResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.DeleteWebhookEndpointRequest{
        EndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.DeleteWebhookEndpoint(
        context.TODO(),
        request,
    )
}
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

**endpointID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.UpdateWebhookEndpoint(EndpointID, request) -> *affinity.UpdateWebhookEndpointResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:write and Idempotency-Key. Updates an endpoint in the selected organization and mode. Omitted practiceIds preserves the filter; an empty array removes the practice filter. Subscription changes apply to newly generated events.
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
request := &affinity.UpdateWebhookEndpointRequest{
        EndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.UpdateWebhookEndpoint(
        context.TODO(),
        request,
    )
}
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

**endpointID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**payloadStyle:** `*affinity.UpdateWebhookEndpointRequestPayloadStyle` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.UpdateWebhookEndpointRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**subscribedEvents:** `[]*affinity.UpdateWebhookEndpointRequestSubscribedEventsItem` 
    
</dd>
</dl>

<dl>
<dd>

**url:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.RotateWebhookEndpointSecret(EndpointID) -> *affinity.RotateWebhookEndpointSecretResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.RotateWebhookEndpointSecretRequest{
        EndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.RotateWebhookEndpointSecret(
        context.TODO(),
        request,
    )
}
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

**endpointID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.TestWebhookEndpoint(EndpointID) -> *affinity.TestWebhookEndpointResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.TestWebhookEndpointRequest{
        EndpointID: "whe_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.TestWebhookEndpoint(
        context.TODO(),
        request,
    )
}
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

**endpointID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListWebhookEvents() -> *affinity.ListWebhookEventsResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListWebhookEventsRequest{
        EndingBefore: affinity.String(
            "evt_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "evt_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Webhooks.ListWebhookEvents(
        context.TODO(),
        request,
    )
}
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

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListWebhookEventsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.GetWebhookEvent(EventID) -> *affinity.GetWebhookEventResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.GetWebhookEventRequest{
        EventID: "evt_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Webhooks.GetWebhookEvent(
        context.TODO(),
        request,
    )
}
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

**eventID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ReplayWebhookEvent(EventID) -> *affinity.ReplayWebhookEventResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ReplayWebhookEventRequest{
        EventID: "evt_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.ReplayWebhookEvent(
        context.TODO(),
        request,
    )
}
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

**eventID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityOrganizationID:** `*string` — Defaults to the API key organization. A platform may select a practice or pharmacy only with an explicit webhook grant in this mode. This changes the webhook owner, not the caller or event subscriptions.
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.ListWebhookGrants() -> *affinity.ListWebhookGrantsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:read on the owning practice or pharmacy key. Lists platform webhook grants in the key's mode. Platforms cannot list or grant themselves delegated access.
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
request := &affinity.ListWebhookGrantsRequest{
        StartingAfter: affinity.String(
            "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Webhooks.ListWebhookGrants(
        context.TODO(),
        request,
    )
}
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

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.SaveWebhookGrant(PlatformID, request) -> *affinity.SaveWebhookGrantResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:write on the owning practice or pharmacy key and Idempotency-Key. Grants or replaces a platform's webhook permissions in this mode. A practice must already be connected to that platform. The grant does not give the platform access to other API resources.
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
request := &affinity.SaveWebhookGrantRequest{
        PlatformID: "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Scopes: []affinity.SaveWebhookGrantRequestScopesItem{
            affinity.SaveWebhookGrantRequestScopesItemWebhooksRead,
        },
    }
client.Webhooks.SaveWebhookGrant(
        context.TODO(),
        request,
    )
}
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

**platformID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*affinity.SaveWebhookGrantRequestScopesItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Webhooks.RevokeWebhookGrant(PlatformID) -> *affinity.RevokeWebhookGrantResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires webhooks:write on the owning practice or pharmacy key and Idempotency-Key. Removes platform webhook access in this mode. Existing endpoints remain owned by the practice or pharmacy and continue operating.
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
request := &affinity.RevokeWebhookGrantRequest{
        PlatformID: "acct_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Webhooks.RevokeWebhookGrant(
        context.TODO(),
        request,
    )
}
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

**platformID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Team
<details><summary><code>client.Team.RegisterUser(PracticeID, request) -> *affinity.RegisterUserResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write and Idempotency-Key. Registers a practice member without an invitation. Test requires synthetic .test emails and Affinity Test NPIs. Live requires approved integration and practice access. Identity attestation records the integration's assertion; it does not verify login email or clinical credentials. Existing memberships and verified provider records are preserved. Use the returned user ID for orders and signing.
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
request := &affinity.RegisterUserRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        ExternalID: "externalId",
        Email: "email",
        Name: "name",
        Role: affinity.RegisterUserRequestRoleAdministrator,
        IdentityAttestation: true,
    }
client.Team.RegisterUser(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*affinity.RegisterUserRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]*affinity.RegisterUserRequestRolesItem` 
    
</dd>
</dl>

<dl>
<dd>

**profileDetails:** `*affinity.RegisterUserRequestProfileDetails` 
    
</dd>
</dl>

<dl>
<dd>

**npi:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**licenses:** `[]*affinity.RegisterUserRequestLicensesItem` 
    
</dd>
</dl>

<dl>
<dd>

**legalName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.RegisterUserRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**locationIDs:** `[]string` 
    
</dd>
</dl>

<dl>
<dd>

**identityAttestation:** `bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.ListPracticeTeamInvitations(PracticeID) -> *affinity.ListPracticeTeamInvitationsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Lists practice invitations, including invitations sent in Clinic. Filter by pending, expired, accepted, declined, or revoked status, exact email, or your integration externalId. Only your integration and API key mode can see its external identity and onboarding state. Follow person.nextActions after invitation acceptance.
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
request := &affinity.ListPracticeTeamInvitationsRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        StartingAfter: affinity.String(
            "invite_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "invite_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Team.ListPracticeTeamInvitations(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPracticeTeamInvitationsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` — Match this integration's external identity in the API key's mode.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.InvitePracticeTeamPerson(PracticeID, request) -> *affinity.InvitePracticeTeamPersonResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write on the practice key or its platform key. Use roles to combine administrator, prescriber, clinical_staff, billing, or developer presets. Ownership uses the protected owner designation. The singular role field remains available for single-role assignments. Creates a real organization invitation and optional prescriber setup. The recipient must accept with their Affinity account. Repeating the same external identity retries pending invitation delivery. Accepted invitations do not change existing access. Team membership is shared between Test and Live; the external identity is mode-scoped. Keys cannot accept invitations. Headless registration and signing use separate endpoints.
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
request := &affinity.InvitePracticeTeamPersonRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        ExternalID: "externalId",
        Email: "email",
        Name: "name",
    }
client.Team.InvitePracticeTeamPerson(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*affinity.InvitePracticeTeamPersonRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]*affinity.InvitePracticeTeamPersonRequestRolesItem` 
    
</dd>
</dl>

<dl>
<dd>

**profileDetails:** `*affinity.InvitePracticeTeamPersonRequestProfileDetails` 
    
</dd>
</dl>

<dl>
<dd>

**npi:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**licenses:** `[]*affinity.InvitePracticeTeamPersonRequestLicensesItem` 
    
</dd>
</dl>

<dl>
<dd>

**legalName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.InvitePracticeTeamPersonRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**locationIDs:** `[]string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetPracticeTeam(PracticeID) -> *affinity.GetPracticeTeamResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Returns counts of members, invitations, and prescribers. Use the paginated members, prescribers, and invitations collections for individual records. Team access and clinician credentials are shared between Test and Live.
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
request := &affinity.GetPracticeTeamRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Team.GetPracticeTeam(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.ListPracticeTeamMembers(PracticeID) -> *affinity.ListPracticeTeamMembersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Search the roster by name or email, and filter by role or membership status. Includes members invited in Clinic, location access, and account-specific prescriber connections. Memberships are shared between Test and Live.
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
request := &affinity.ListPracticeTeamMembersRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        StartingAfter: affinity.String(
            "mbr_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "mbr_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Team.ListPracticeTeamMembers(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*affinity.ListPracticeTeamMembersRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPracticeTeamMembersRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.ListPracticeTeamPrescribers(PracticeID) -> *affinity.ListPracticeTeamPrescribersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Filter practice prescribers by name, NPI, state, and practice status. Records include submitted licenses and their IDs. Signing authority also requires an active account connection, Live practice access, and prescription eligibility.
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
request := &affinity.ListPracticeTeamPrescribersRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        StartingAfter: affinity.String(
            "prov_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "prov_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Team.ListPracticeTeamPrescribers(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**npi:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` — Match a submitted license jurisdiction. This does not establish signing eligibility.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPracticeTeamPrescribersRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetPracticeTeamMember(PracticeID, MemberID) -> *affinity.GetPracticeTeamMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Returns current account membership, roles, location access, and prescriber connection. The member ID identifies practice access; it is not the integration user ID used by orders.
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
request := &affinity.GetPracticeTeamMemberRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        MemberID: "mbr_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Team.GetPracticeTeamMember(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.UpdatePracticeTeamMember(PracticeID, MemberID, request) -> *affinity.UpdatePracticeTeamMemberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write. Supply role, status, or locationIds; omitted values stay unchanged. A role replaces existing roles. Disable access with status disabled. An empty locationIds array grants all practice locations. Ownership changes require an active practice owner using a personal API key; service keys manage non-owner memberships. The final active owner cannot be removed. Changes apply to both Test and Live. Sign-in email and account security remain account settings.
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
request := &affinity.UpdatePracticeTeamMemberRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        MemberID: "mbr_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Team.UpdatePracticeTeamMember(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**memberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**role:** `*affinity.UpdatePracticeTeamMemberRequestRole` 
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]*affinity.UpdatePracticeTeamMemberRequestRolesItem` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.UpdatePracticeTeamMemberRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**locationIDs:** `[]string` — Replace location access. An empty array grants access to all practice locations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetPracticeTeamPrescriber(PracticeID, PrescriberID) -> *affinity.GetPracticeTeamPrescriberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Returns the clinical profile and submitted licenses, including license IDs. This is setup information, not a signing authorization.
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
request := &affinity.GetPracticeTeamPrescriberRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PrescriberID: "prov_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Team.GetPracticeTeamPrescriber(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriberID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.UpdatePracticeTeamPrescriber(PracticeID, PrescriberID, request) -> *affinity.UpdatePracticeTeamPrescriberResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write. Set practiceStatus to inactive to remove prescribing access in this practice, or active to restore an existing association. This does not create membership or signing authority. Practice status applies to Test and Live. Shared identity and license edits require Affinity support.
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
request := &affinity.UpdatePracticeTeamPrescriberRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PrescriberID: "prov_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Team.UpdatePracticeTeamPrescriber(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**legalName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.UpdatePracticeTeamPrescriberRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**practiceStatus:** `*affinity.UpdatePracticeTeamPrescriberRequestPracticeStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.CreatePracticeTeamLicense(PracticeID, PrescriberID, request) -> *affinity.CreatePracticeTeamLicenseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write and an active accepted prescriber account connection in this practice. Adds a license. Expiration is optional, but must be in the future when supplied. An exact repeat returns the existing license; update an existing license with PATCH and its license ID. Licenses are shared across practices and Test/Live. Other licenses stay unchanged.
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
request := &affinity.CreatePracticeTeamLicenseRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PrescriberID: "prov_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        State: "state",
        LicenseNumber: "licenseNumber",
    }
client.Team.CreatePracticeTeamLicense(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**licenseNumber:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.UpdatePracticeTeamLicense(PracticeID, PrescriberID, LicenseID, request) -> *affinity.UpdatePracticeTeamLicenseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write and an active accepted prescriber account connection in this practice. Correct the state or license number, or set or clear the optional expiresAt value. A supplied expiration must be in the future. Other licenses stay unchanged. Changes apply across practices and Test/Live.
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
request := &affinity.UpdatePracticeTeamLicenseRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PrescriberID: "prov_01j2y8m6jcc9tt24af5pw9x1bc",
        LicenseID: "lic_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Team.UpdatePracticeTeamLicense(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescriberID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**licenseID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**state:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**licenseNumber:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.GetPracticeTeamInvitation(PracticeID, InvitationID) -> *affinity.GetPracticeTeamInvitationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:read. Returns invitation status and current onboarding state for your integration. An accepted invitation can still have disabled membership or pending clinical review. Invitation tokens are never returned.
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
request := &affinity.GetPracticeTeamInvitationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        InvitationID: "invite_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Team.GetPracticeTeamInvitation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.RevokePracticeTeamInvitation(PracticeID, InvitationID) -> *affinity.RevokePracticeTeamInvitationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write. Revokes a pending or expired invitation and its pending prescriber account connection. Repeating the revoke returns the revoked invitation. Accepted invitations return 409; disable the member instead. Retains invitation history.
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
request := &affinity.RevokePracticeTeamInvitationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        InvitationID: "invite_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Team.RevokePracticeTeamInvitation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Team.ResendPracticeTeamInvitation(PracticeID, InvitationID) -> *affinity.ResendPracticeTeamInvitationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires team:write. Resends a pending or expired invitation with the same ID, recipient, roles, and locations. The previous link stops working and the new link expires in seven days. Accepted and revoked invitations return 409. A 502 means the invitation was saved but email delivery could not be confirmed; retry this operation.
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
request := &affinity.ResendPracticeTeamInvitationRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        InvitationID: "invite_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Team.ResendPracticeTeamInvitation(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Patients
<details><summary><code>client.Patients.ListPatientAddresses(PracticeID, PatientID) -> *affinity.ListPatientAddressesResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.ListPatientAddressesRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        StartingAfter: affinity.String(
            "addr_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        EndingBefore: affinity.String(
            "addr_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Patients.ListPatientAddresses(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPatientAddressesRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.CreatePatientAddress(PracticeID, PatientID, request) -> *affinity.CreatePatientAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the existing active address for a normalized duplicate. The first address becomes the default. API keys require Idempotency-Key.
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
request := &affinity.CreatePatientAddressRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Address: &affinity.CreatePatientAddressRequestAddress{
            City: "city",
            Line1: "line1",
            PostalCode: "postalCode",
            State: "state",
        },
    }
client.Patients.CreatePatientAddress(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.CreatePatientAddressRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**preferredShipping:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**recipientName:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.ArchivePatientAddress(PracticeID, PatientID, AddressID) -> *affinity.ArchivePatientAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Preserves the address ID and history. Archiving the default selects the oldest remaining active address. Existing orders remain unchanged.
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
request := &affinity.ArchivePatientAddressRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        AddressID: "addr_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Patients.ArchivePatientAddress(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**addressID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.UpdatePatientAddress(PracticeID, PatientID, AddressID, request) -> *affinity.UpdatePatientAddressResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &affinity.UpdatePatientAddressRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        AddressID: "addr_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Patients.UpdatePatientAddress(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**addressID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.UpdatePatientAddressRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**label:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**recipientName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**preferredShipping:** `*bool` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.SetDefaultPatientAddress(PracticeID, PatientID, AddressID) -> *affinity.SetDefaultPatientAddressResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Changes delivery selection for future drafts, without changing patient clinical location or existing signed orders.
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
request := &affinity.SetDefaultPatientAddressRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        AddressID: "addr_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Patients.SetDefaultPatientAddress(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**addressID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.ListPatients(PracticeID) -> *affinity.ListPatientsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists patients in one practice and mode. Use externalId for an exact match in the calling integration's namespace. Use externalIdentitySource with externalIdentityValue to search an explicit alias. Identity matching is case-sensitive after trimming whitespace. Other filters also apply.
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
request := &affinity.ListPatientsRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        EndingBefore: affinity.String(
            "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Patients.ListPatients(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalIdentitySource:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalIdentityValue:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**gender:** `*affinity.ListPatientsRequestGender` 
    
</dd>
</dl>

<dl>
<dd>

**lastOrderAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastOrderBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**program:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**query:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*affinity.ListPatientsRequestSort` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**states:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.ListPatientsRequestStatus` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.CreatePatient(PracticeID, request) -> *affinity.CreatePatientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a patient or resolves a matching externalId or external identity within this practice and mode. externalId belongs to the calling integration; externalIdentities holds aliases from other systems. Resolution preserves existing demographics; use PATCH to update them. Conflicting identifiers return 409. Email never merges patients. API keys require Idempotency-Key.
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
request := &affinity.CreatePatientRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        DateOfBirth: "dateOfBirth",
        Name: &affinity.CreatePatientRequestName{
            First: "first",
            Last: "last",
        },
    }
client.Patients.CreatePatient(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.CreatePatientRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**clinicalProfile:** `*affinity.CreatePatientRequestClinicalProfile` 
    
</dd>
</dl>

<dl>
<dd>

**dateOfBirth:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalIdentities:** `[]*affinity.CreatePatientRequestExternalIdentitiesItem` 
    
</dd>
</dl>

<dl>
<dd>

**addresses:** `[]*affinity.CreatePatientRequestAddressesItem` 
    
</dd>
</dl>

<dl>
<dd>

**encounters:** `[]*affinity.CreatePatientRequestEncountersItem` 
    
</dd>
</dl>

<dl>
<dd>

**gender:** `*affinity.CreatePatientRequestGender` 
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**medicalRecordNumber:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**measurements:** `[]*affinity.CreatePatientRequestMeasurementsItem` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*affinity.CreatePatientRequestName` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**programs:** `[]*affinity.CreatePatientRequestProgramsItem` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.GetPatient(PracticeID, PatientID) -> *affinity.GetPatientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns one patient in the authorized practice and mode.
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
request := &affinity.GetPatientRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Patients.GetPatient(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.DeletePatient(PracticeID, PatientID) -> *affinity.DeletePatientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires patients:write and Idempotency-Key for API keys. Permanently deletes a patient with no order history. Any order history returns 409; use Update patient with status archived instead. Available to practice keys and authorized platform keys. Reusing the same idempotency key returns the original deletion result.
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
request := &affinity.DeletePatientRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Patients.DeletePatient(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.UpdatePatient(PracticeID, PatientID, request) -> *affinity.UpdatePatientResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a patient in the current practice and mode. Omitted fields remain unchanged; null clears an optional field. externalId updates the calling integration's identifier. externalIdentities replaces its explicit aliases. Identifiers cannot be reassigned from another patient. API keys require Idempotency-Key.
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
request := &affinity.UpdatePatientRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
    }
client.Patients.UpdatePatient(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.UpdatePatientRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**clinicalProfile:** `*affinity.UpdatePatientRequestClinicalProfile` 
    
</dd>
</dl>

<dl>
<dd>

**dateOfBirth:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**externalIdentities:** `[]*affinity.UpdatePatientRequestExternalIdentitiesItem` 
    
</dd>
</dl>

<dl>
<dd>

**addresses:** `[]*affinity.UpdatePatientRequestAddressesItem` 
    
</dd>
</dl>

<dl>
<dd>

**encounters:** `[]*affinity.UpdatePatientRequestEncountersItem` 
    
</dd>
</dl>

<dl>
<dd>

**gender:** `*affinity.UpdatePatientRequestGender` 
    
</dd>
</dl>

<dl>
<dd>

**locationID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**medicalRecordNumber:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**measurements:** `[]*affinity.UpdatePatientRequestMeasurementsItem` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*affinity.UpdatePatientRequestName` 
    
</dd>
</dl>

<dl>
<dd>

**programs:** `[]*affinity.UpdatePatientRequestProgramsItem` 
    
</dd>
</dl>

<dl>
<dd>

**phone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*affinity.UpdatePatientRequestStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.GetPatientAllergies(PracticeID, PatientID) -> *affinity.GetPatientAllergiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the patient's structured allergy entries and review status. A not_reviewed status is not a no-known-allergies assertion and blocks clinical review and signing.
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
request := &affinity.GetPatientAllergiesRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Patients.GetPatientAllergies(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Patients.ReplacePatientAllergies(PracticeID, PatientID, request) -> *affinity.ReplacePatientAllergiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replaces the patient's structured allergy record. Sending no_known is the explicit no-known-allergies acknowledgement; recorded requires at least one entry. Idempotency-Key is required.
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
request := &affinity.ReplacePatientAllergiesRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        PatientID: "pat_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        Allergies: []*affinity.ReplacePatientAllergiesRequestAllergiesItem{
            &affinity.ReplacePatientAllergiesRequestAllergiesItem{
                Category: affinity.ReplacePatientAllergiesRequestAllergiesItemCategoryDrug,
                Reactions: []*affinity.ReplacePatientAllergiesRequestAllergiesItemReactionsItem{
                    &affinity.ReplacePatientAllergiesRequestAllergiesItemReactionsItem{
                        Display: "display",
                    },
                },
                Source: affinity.ReplacePatientAllergiesRequestAllergiesItemSourceDoctor,
                Substance: "substance",
                VerificationStatus: affinity.ReplacePatientAllergiesRequestAllergiesItemVerificationStatusUnconfirmed,
            },
        },
        ReviewStatus: affinity.ReplacePatientAllergiesRequestReviewStatusNotReviewed,
    }
client.Patients.ReplacePatientAllergies(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**patientID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**affinityActorID:** `*string` — Required for user actors and optional for system actors. Omit both actor headers to use the authenticated service account as a system actor.
    
</dd>
</dl>

<dl>
<dd>

**affinityActorType:** `*string` — Use user when a person initiated the action and system for autonomous work. Omit both actor headers to default to system.
    
</dd>
</dl>

<dl>
<dd>

**allergies:** `[]*affinity.ReplacePatientAllergiesRequestAllergiesItem` 
    
</dd>
</dl>

<dl>
<dd>

**reviewStatus:** `*affinity.ReplacePatientAllergiesRequestReviewStatus` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Practices
<details><summary><code>client.Practices.ListPractices() -> *affinity.ListPracticesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the practices that belong to the platform. The default Affinity-Version is 2026-09-28.
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
request := &affinity.ListPracticesRequest{
        EndingBefore: affinity.String(
            "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
        StartingAfter: affinity.String(
            "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.Practices.ListPractices(
        context.TODO(),
        request,
    )
}
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

**search:** `*string` — Case-insensitive search by practice name or external ID.
    
</dd>
</dl>

<dl>
<dd>

**endingBefore:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**startingAfter:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.CreatePractice(request) -> *affinity.CreatePracticeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a practice owned by the platform. Set liveEnabled to true to enable Live access at creation with an approved platform and a Live request. Defaults to false. Requires practices:write. Send Idempotency-Key when you retry the same request.
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
request := &affinity.CreatePracticeRequest{
        Address: &affinity.CreatePracticeRequestAddress{
            City: "Los Angeles",
            Country: affinity.String(
                "US",
            ),
            Line1: "100 Practice Way",
            PostalCode: "90001",
            State: "CA",
        },
        Attestations: &affinity.CreatePracticeRequestAttestations{
            AuthorizedPracticeRelationship: true,
            AuthorizedPhiTransfer: true,
            MinimumNecessaryPhi: true,
            ProviderDataAccuracy: true,
        },
        ExternalID: affinity.String(
            "practice_123",
        ),
        LegalName: affinity.String(
            "Example Medical Group PLLC",
        ),
        Metadata: map[string]any{
            "key": "value",
        },
        Name: "Example Medical Group",
        Prescribers: []*affinity.CreatePracticeRequestPrescribersItem{
            &affinity.CreatePracticeRequestPrescribersItem{
                Credentials: affinity.String(
                    "MD",
                ),
                LicenseStates: []string{
                    "CA",
                },
                Name: "Alex Morgan",
                Npi: "1234567893",
            },
        },
        PrimaryContact: &affinity.CreatePracticeRequestPrimaryContact{
            Email: "operations@example-practice.com",
            Name: "Jordan Lee",
        },
        SupportEmail: affinity.String(
            "support@example-practice.com",
        ),
    }
client.Practices.CreatePractice(
        context.TODO(),
        request,
    )
}
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

**idempotencyKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**liveEnabled:** `*bool` — Enable Live access at creation. Requires an approved platform and a Live request. Defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.CreatePracticeRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**attestations:** `*affinity.CreatePracticeRequestAttestations` 
    
</dd>
</dl>

<dl>
<dd>

**complianceContact:** `*affinity.CreatePracticeRequestComplianceContact` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**legalName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**prescribers:** `[]*affinity.CreatePracticeRequestPrescribersItem` 
    
</dd>
</dl>

<dl>
<dd>

**primaryContact:** `*affinity.CreatePracticeRequestPrimaryContact` 
    
</dd>
</dl>

<dl>
<dd>

**supportEmail:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**supportPhone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Optional IANA timezone override. Omit to leave unchanged; null clears it. No timezone is inferred when creating a record.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.GetPractice(PracticeID) -> *affinity.GetPracticeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns one practice that belongs to the platform.
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
request := &affinity.GetPracticeRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Practices.GetPractice(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Practices.UpdatePractice(PracticeID, request) -> *affinity.UpdatePracticeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates one practice owned by the platform. Set liveEnabled to true or false to control Live access with an approved platform and a Live request. Affinity Admin decisions take precedence. Requires practices:write. Send Idempotency-Key when you retry the same request.
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
request := &affinity.UpdatePracticeRequest{
        PracticeID: "prac_01j2y8m6jcc9tt24af5pw9x1bc",
    }
client.Practices.UpdatePractice(
        context.TODO(),
        request,
    )
}
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

**practiceID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**liveEnabled:** `*bool` — Enable or disable Live access for an owned practice. Requires an approved platform and a Live request. Affinity Admin decisions take precedence.
    
</dd>
</dl>

<dl>
<dd>

**address:** `*affinity.UpdatePracticeRequestAddress` 
    
</dd>
</dl>

<dl>
<dd>

**attestations:** `*affinity.UpdatePracticeRequestAttestations` 
    
</dd>
</dl>

<dl>
<dd>

**complianceContact:** `*affinity.UpdatePracticeRequestComplianceContact` 
    
</dd>
</dl>

<dl>
<dd>

**externalID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**legalName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `map[string]any` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**prescribers:** `[]*affinity.UpdatePracticeRequestPrescribersItem` 
    
</dd>
</dl>

<dl>
<dd>

**primaryContact:** `*affinity.UpdatePracticeRequestPrimaryContact` 
    
</dd>
</dl>

<dl>
<dd>

**supportEmail:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**supportPhone:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**timezone:** `*string` — Optional IANA timezone override. Omit to leave unchanged; null clears it. No timezone is inferred when creating a record.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Platform Pricing
<details><summary><code>client.PlatformPricing.PlatformPublicAPISellingPricesReadSellingPrice(CatalogItemID) -> *affinity.PlatformPublicAPISellingPricesReadSellingPriceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires selling_prices:read. Omit practiceId for the platform default, or supply a managed practice. A null amount inherits the next applicable price. Amounts use the catalog pricing basis, in USD cents. purchaseAmountCents is the platform's Affinity purchase price for that same basis. requiresReview indicates changed product pricing terms, not a below-purchase-price discount.
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
request := &affinity.PlatformPublicAPISellingPricesReadSellingPriceRequest{
        CatalogItemID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        PracticeID: affinity.String(
            "prac_01j2y8m6jcc9tt24af5pw9x1bc",
        ),
    }
client.PlatformPricing.PlatformPublicAPISellingPricesReadSellingPrice(
        context.TODO(),
        request,
    )
}
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

**catalogItemID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.PlatformPricing.PlatformPublicAPISellingPricesUpdateSellingPrice(CatalogItemID, request) -> *affinity.PlatformPublicAPISellingPricesUpdateSellingPriceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Requires selling_prices:write. Sets a platform default or managed practice override in the current Test/Live mode. Send baseVersion from Read selling price. Null removes the override. Prices use the catalog pricing basis. Intentional discounts below purchaseAmountCents are allowed; compare these amounts to warn about selling below your Affinity purchase price. This does not change the platform's Affinity purchase price or collect practice payments.
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
request := &affinity.PlatformPublicAPISellingPricesUpdateSellingPriceRequest{
        CatalogItemID: "cat_01j2y8m6jcc9tt24af5pw9x1bc",
        IdempotencyKey: "Idempotency-Key",
        BaseVersion: 1,
    }
client.PlatformPricing.PlatformPublicAPISellingPricesUpdateSellingPrice(
        context.TODO(),
        request,
    )
}
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

**catalogItemID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**idempotencyKey:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**practiceID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**amountCents:** `*int` 
    
</dd>
</dl>

<dl>
<dd>

**baseVersion:** `int` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

