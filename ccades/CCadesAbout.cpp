#include "stdafx.h"
#include <stdlib.h>
#include "CCadesAbout.h"
#include "CPPCadesAbout.h"

using namespace CryptoPro::PKI::CAdES;

struct CCadesAbout_t
{
    void *obj;
};

CCadesAbout *CCadesAbout_create()
{
    CCadesAbout *m;
    CPPCadesAboutObject *obj;

    m = (typeof(m))malloc(sizeof(*m));
    obj = new CPPCadesAboutObject();
    m->obj = obj;

    return m;
}

void CCadesAbout_destroy(CCadesAbout *m)
{
    if (m == NULL)
        return;
    delete static_cast<CPPCadesAboutObject *>(m->obj);
    free(m);
}

char* CCadesAbout_get_version(CCadesAbout_t *m)
{
    CPPCadesAboutObject *obj;

    if (m == NULL)
        return 0;

    obj = static_cast<CPPCadesAboutObject *>(m->obj);
    CAtlString sValue;
    obj->get_Version(sValue);
    return (char*)sValue.GetString();
}

int CCadesAbout_get_major_version(CCadesAbout_t *m)
{
    CPPCadesAboutObject *obj;

    if (m == NULL)
        return 0;

    obj = static_cast<CPPCadesAboutObject *>(m->obj);
    DWORD r;
    obj->get_MajorVersion(&r);
    return r;
}

int CCadesAbout_get_minor_version(CCadesAbout_t *m)
{
    CPPCadesAboutObject *obj;

    if (m == NULL)
        return 0;

    obj = static_cast<CPPCadesAboutObject *>(m->obj);
    DWORD r;
    obj->get_MinorVersion(&r);
    return r;
}

int CCadesAbout_get_build_version(CCadesAbout_t *m)
{
    CPPCadesAboutObject *obj;

    if (m == NULL)
        return 0;

    obj = static_cast<CPPCadesAboutObject *>(m->obj);
    DWORD r;
    obj->get_BuildVersion(&r);
    return r;
}
