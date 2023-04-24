This Project has below services

- Customer:
   - Provides information about already available customers which can be used for testing.

- Product:
   - Provides information about already existing products with details like name, quantity, price, category which can be used for testing.
   - Provides service to add the new product with id starting with **p**.

- Order:
   - Provides information about the order like dispatchDate, orderStatus, orderedItems, total amount, discount amount.
   - The user can get the product from the catalogue and using that info should be able to place an order.
   - Once the order is placed for a particular product, the product catalogue should be updated accordingly.
     (Max quantity of a particular product that can be ordered is 10)
   - If the order contains 3 premium different products, order value will be discounted by 10%
   - The Order service will be able to update the orderStatus for a particular order.DispatchDate for and order will be populated only when the orderStatus is 'Dispatched'.

Config Values : config.env file is placed under ./config folder with below data
   - DISCOUNT=10
   - DISCOUNT_PRODUCTS_CATEGORY=Premium
   - DISCOUNT_PRODUCTS_COUNT=3
   - MAX_PRODUCTS_IN_ORDER=10

Possible Values For few fields are as below. Used gin package **binding** tag to validate these fields.
   - product category values: Premium/Regular/Budget
   - order status values: Placed/Dispatched/Completed/Cancelled

To run this Project go to ./cmd folder and run **go run main.go**

For endpoints testing, Postman collection is placed under ./docs folder

**Unit testing**
1. Added testcases under handler and service, wherever required.
2. Used Mock to make sure, it doesn't affect the datastore.
3. Run **go test -cover ./...** under this current directory to see the coverage of all packages.

**NOTE**
1. For Ease of usage, few products and customers are already added.
2. Datastore used is in-memory map instead of actual db.
