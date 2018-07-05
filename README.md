# CostumeQuoterFlask

Same "Material" API for ArtistCity CostumeQuoter, using Go and MySql (I think this will be the good one).

## Install requirements

dep ensure

## Run

go run main.go

## API

### **getMaterials**

Returns all Materials from collection Materials

- **URL**

  /material

- **Method:**

  `GET`

### **createMaterial**

Create new material based on JSON request

- **URL**

  /material

- **Method:**

  `POST`

### **getMaterialByObjectId**

Returns one Material by ObjectId

- **URL**

  /material/<id>

- **Method:**

  `GET`
