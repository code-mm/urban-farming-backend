# Endpoints


## Path /farm

### Path /
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: NONE
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns details of logged in device.

### Path /datapoint/ph
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all ph values for the logged in device or takes new ph values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).

### Path /datapoint/oxygen
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all oxygen values for the logged in device or takes new oxygen values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).

### Path /datapoint/temperature
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all temperature values for the logged in device or takes new temperature values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).


## Path /authentication

### Path /gettoken/device
Allowed HTTP method: POST
Allowed content type: application/x-www-form-urlencoded
Required fields: identifier, secret
Optional fields: NONE
Return status type: 200 on success or 401 if login credentials are wrong
Return content type: application/jwt
Description: Takes device login credentials as form data and returns JSON web token on success.

### Path /gettoken/user
Allowed HTTP method: POST
Allowed content type: application/x-www-form-urlencoded
Required fields: email, password
Optional fields: NONE
Return status type: 200 on success or 401 if login credentials are wrong
Return content type: application/jwt
Description: Takes device login credentials as form data and returns JSON web token on success.


## Path /user

### Path /
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: NONE
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns details of logged in user.

### Path /farm/
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: NONE
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns list of available farms for logged in user.

### Path /farm/<farmId>/datapoint/ph
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: farmId
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns datapoints ph for given farm if user has access permissions.

### Path /farm/<farmId>/datapoint/oxygen
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: farmId
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns datapoints ph for given farm if user has access permissions.

### Path /farm/<farmId>/datapoint/temperature
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: farmId
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns datapoints ph for given farm if user has access permissions.
