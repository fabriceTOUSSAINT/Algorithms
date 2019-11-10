# How Apollo, Graphql, React Hooks and refactoring a failed project made me a better developer

> This is within a timeline of roughly 3-weeks.

### Technologies working with
* Apollo
* Graphql
* React (functional)
* Postgresql
* API Design
* Node/Express

---

## Learning Apollo & Graphql

### What is Apollo?
* Framework for implementing GraphQL
* Provides the data graph layer that connects modern apps to the cloud
* Data is organized under on connected data graph
  * you can browse everything available
  * join data across multiple sources
  * get results in the shap you need and on any platform
* Suited for Microservices




### Part 1: Schemas
* Blueprint of all the data I can access in my graph
* defines what data we can fetch through queries
* defines what data we can update through mutations
* strongly typed

> When Developing in Apollo, practice **Schema First Development**, essentially don't begin work until we design a Schema around the needs of the clints that are using them.
>Think about the data we will need to expose in order to build our app

> ### Lens_Search //TODO: expand Schema from this?
> ---
> * Fetch list of lens ( used for searching against/autofill)
> * Fetch Lens by its ID
> * Fetch for batch of photos shot with lens (by ID)
> * Login User
> * Save lens to favorites if logged in
> * Fetch list of saved favorites if logged in


`...Continue`
#### Query type
* our **Entry Point** into our schema
* describes **What** data we can fetch
* We use GraphQL's language for schema definition (SDL)
* A _Special object type_

#### Mutation type
* The **Entry Point** into our graph for modifying data
* A _Special object type_
*

> Types fields can contain arguments, not just queries

* Just setting up the Schema makes everything much clearer, Being able to spin up an Apollo server as light and quickly as Express, define a Schema and you have an endpoint! amazing!


### Part 2: Hooking up Data sources
* Connect to either REST or own Database
* // I think? it simply extends a "DataSource" class, I believe they act as endpoints


### Part 3: Writing my graph's resolvers
* Resolvers provide the instructions for turning a GraphQL operation ( a query, mutation, or subscription) into data
* Either return the sam type of data specefied in our schema
* or
* a promise for that data

```
// Resolver Syntax
/**
 * parent - object contains the result returned from the resolver on the parent type
 * args - object the contains the arguments passed to the field
 * context - object shared by all resolvers in a GraphQL operation.
 *           used to contain per-request state suc has authentication information and access to our data sources
 * info - information about the execution state of the operation which should only
 be used in advanced cases
 */
fieldName: (parent, args, context, info) => data;
```
