/**
 * BSpec Rust SDK C Interface
 *
 * C-compatible header for integrating BSpec functionality with
 * Go CLI tools, native macOS apps, and other languages.
 */

#ifndef BSPEC_H
#define BSPEC_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>
#include <stdint.h>

// Opaque types
typedef struct Archive Archive;
typedef struct QueryResult QueryResult;

// Error codes
typedef enum {
    BSPEC_SUCCESS = 0,
    BSPEC_INVALID_PATH = 1,
    BSPEC_FILE_NOT_FOUND = 2,
    BSPEC_INVALID_ARCHIVE = 3,
    BSPEC_INVALID_QUERY = 4,
    BSPEC_SERIALIZATION_ERROR = 5,
    BSPEC_VALIDATION_ERROR = 6,
    BSPEC_UNKNOWN_ERROR = 99
} BSpecErrorCode;

// Archive operations
/**
 * Open a .bspec archive file
 * @param path Path to the .bspec file
 * @return Archive pointer or NULL on error
 */
Archive* bspec_open_archive(const char* path);

/**
 * Create a new empty archive
 * @param name Name for the new archive
 * @return Archive pointer or NULL on error
 */
Archive* bspec_create_archive(const char* name);

/**
 * Save archive to file
 * @param archive Archive to save
 * @param path Output file path
 * @return Error code
 */
BSpecErrorCode bspec_save_archive(Archive* archive, const char* path);

/**
 * Close and free an archive
 * @param archive Archive to close
 */
void bspec_close_archive(Archive* archive);

// Query operations
/**
 * Query documents in archive
 * @param archive Archive to query
 * @param query_json JSON query string
 * @return QueryResult pointer or NULL on error
 */
QueryResult* bspec_query_documents(Archive* archive, const char* query_json);

/**
 * Get the number of results in a query
 * @param result Query result
 * @return Number of documents found
 */
int bspec_query_result_count(QueryResult* result);

/**
 * Get a document JSON from query results by index
 * @param result Query result
 * @param index Document index
 * @return JSON string (must be freed with bspec_free_string)
 */
const char* bspec_get_document_json(QueryResult* result, size_t index);

/**
 * Free query results
 * @param result Query result to free
 */
void bspec_free_query_result(QueryResult* result);

// Utility operations
/**
 * Get archive statistics as JSON
 * @param archive Archive
 * @return JSON statistics string (must be freed with bspec_free_string)
 */
const char* bspec_get_archive_stats(Archive* archive);

/**
 * Free a string returned by BSpec functions
 * @param str String to free
 */
void bspec_free_string(char* str);

// Document validation
/**
 * Validate a document JSON string
 * @param document_json Document JSON
 * @return Error code
 */
BSpecErrorCode bspec_validate_document(const char* document_json);

// Archive manipulation
/**
 * Add document to archive
 * @param archive Archive
 * @param document_json Document JSON
 * @return Error code
 */
BSpecErrorCode bspec_add_document(Archive* archive, const char* document_json);

/**
 * Remove document from archive
 * @param archive Archive
 * @param document_id Document ID
 * @return Error code
 */
BSpecErrorCode bspec_remove_document(Archive* archive, const char* document_id);

/**
 * Add asset to archive
 * @param archive Archive
 * @param path Asset path within archive
 * @param data Asset data
 * @param data_len Length of asset data
 * @return Error code
 */
BSpecErrorCode bspec_add_asset(Archive* archive, const char* path, const uint8_t* data, size_t data_len);

// Error handling
/**
 * Get error message for error code
 * @param error_code Error code
 * @return Error message string (do not free)
 */
const char* bspec_get_error_message(BSpecErrorCode error_code);

// Version information
/**
 * Get BSpec SDK version
 * @return Version string (do not free)
 */
const char* bspec_get_version(void);

/**
 * Get SDK build information
 * @return Build info string (do not free)
 */
const char* bspec_get_build_info(void);

#ifdef __cplusplus
}
#endif

#endif /* BSPEC_H */