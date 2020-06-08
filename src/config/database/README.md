# database package

This package is responsible for establishing a connection to the database and applying the applications's database migrations. 

When a connection attempt is made, a retry mechanism will attempt to successully ping the database for a total of 30 retries. If no successful ping is made, the application panic closes.

The MigrationRunner struct is responsible for applying the database migrations stored in the [migrations](../../../resources/db/migration) folder. This component automatically creates a `public.database_version` table where migration checksums are stored.