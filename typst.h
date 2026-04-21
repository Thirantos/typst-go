#pragma once

#include <stdint.h>

typedef struct TypstDocument TypstDocument;

typedef struct TypstWorld TypstWorld;

typedef struct TypstError {
  char *message;
} TypstError;

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

struct TypstWorld *typst_world_new(char *root, char *source);

struct TypstDocument *typst_world_compile(struct TypstWorld *world,
                                          struct TypstError *err);

void typst_document_to_pdf(struct TypstDocument *document, uintptr_t *len,
                           uint8_t **data, struct TypstError *err);

void typst_world_free(struct TypstWorld *ptr);

void typst_document_free(struct TypstDocument *ptr);

void typst_pdf_free(uint8_t *ptr, uintptr_t len);

#ifdef __cplusplus
} // extern "C"
#endif // __cplusplus
