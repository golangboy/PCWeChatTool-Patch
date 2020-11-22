#include <Windows.h>
#include <stdlib.h>
#include <string.h>
#include "./_obj/_cgo_export.h"
typedef void (*pSetCallBack)(DWORD func);

char *wchar2char(const wchar_t *wstr)
{
	size_t len = wcslen(wstr) + 1;
	char *cstr = 0;
	size_t converted = 0;
	DWORD iSize = WideCharToMultiByte(CP_ACP, 0, wstr, -1, NULL, 0, NULL, NULL); //iSize =wcslen(pwsUnicode)+1=6
	cstr = (char *)malloc(iSize * sizeof(char));								 //不需要 pszMultiByte = (char*)malloc(iSize*sizeof(char)+1);
	WideCharToMultiByte(CP_ACP, 0, wstr, -1, cstr, iSize, NULL, NULL);
	return cstr;
}
wchar_t *char2wchar(const char *cchar)
{
	wchar_t *m_wchar;
	int len = MultiByteToWideChar(CP_ACP, 0, cchar, strlen(cchar), NULL, 0);
	m_wchar = malloc(sizeof(wchar_t) * (len + 1));
	MultiByteToWideChar(CP_ACP, 0, cchar, strlen(cchar), m_wchar, len);
	m_wchar[len] = '\0';
	return m_wchar;
}
void __attribute__((__stdcall__)) textMessage(const wchar_t *wxId, const wchar_t *msg, const int type)
{
	char *chWxId = wchar2char(wxId);
	char *chMsg = wchar2char(msg);
	//MessageBoxA(NULL, chWxId, chMsg, MB_OK);
	GoString p1, p2;
	p1.p = chWxId;
	p1.n = strlen(chWxId);

	p2.p = chMsg;
	p2.n = strlen(chMsg);
	recvMsg(p1, p2, type);

	free(chWxId);
	free(chMsg);
}

void __attribute__((__stdcall__)) imageMessage(const wchar_t *wxId, const wchar_t *imgPath)
{
	char *chWxId = wchar2char(wxId);
	char *chImgPath = wchar2char(imgPath);
	//MessageBoxA(NULL, chWxId, chMsg, MB_OK);
	GoString p1, p2;
	p1.p = chWxId;
	p1.n = strlen(chWxId);

	p2.p = chImgPath;
	p2.n = strlen(chImgPath);
	recvImageMsg(p1, p2);

	free(chWxId);
	free(chImgPath);
}

void __attribute__((__stdcall__)) voiceMessage(const wchar_t *wxId, const char *ptrData, int dataLen)
{
	char *chWxId = wchar2char(wxId);

	GoString p1;
	p1.p = chWxId;
	p1.n = strlen(chWxId);

	recvVoiceMsg(p1, ptrData, dataLen);

	free(chWxId);
}

void SendTextMessage(char *wxId, char *textMsg)
{
	HMODULE hModule = LoadLibrary("Tool.dll");
	if (0 != hModule)
	{
		typedef(__attribute__((__stdcall__)) * ptrSendText)(wchar_t * wxId, wchar_t * texMsg);
		ptrSendText sendText = GetProcAddress(hModule, "Call_SendTextMessage");
		if (sendText != 0)
		{
			wchar_t *wsId = char2wchar(wxId);
			wchar_t *wsTextMsg = char2wchar(textMsg);

			sendText(wsId, wsTextMsg);

			free(wsTextMsg);
			free(wsId);
		}
	}
}
void SendFileMessage(char *wxId, char *filePath)
{
	HMODULE hModule = LoadLibrary("Tool.dll");
	if (0 != hModule)
	{
		typedef(__attribute__((__stdcall__)) * ptrSendFile)(wchar_t * wxId, wchar_t * filePath);
		ptrSendFile sendFile = GetProcAddress(hModule, "Call_SendFileMessage");
		if (sendFile != 0)
		{
			wchar_t *wsId = char2wchar(wxId);
			wchar_t *wsFilePath = char2wchar(filePath);

			sendFile(wsId,wsFilePath);

			free(wsId);
			free(wsFilePath);
		}
	}
}
void StartHook()
{
	HMODULE hModule = LoadLibrary("Tool.dll");
	if (0 != hModule)
	{
		pSetCallBack addr = GetProcAddress(hModule, "setCallBack_TextMessage");
		if (addr != NULL)
		{
			addr(textMessage);
		}

		addr = GetProcAddress(hModule, "setCallBack_ImageMessage");
		if (addr != NULL)
		{
			addr(imageMessage);
		}
		addr = GetProcAddress(hModule, "setCallBack_VoiceMessage");
		if (addr != NULL)
		{
			addr(voiceMessage);
		}
	}
}