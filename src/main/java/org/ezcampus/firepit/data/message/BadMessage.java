package org.ezcampus.firepit.data.message;

import com.fasterxml.jackson.annotation.JsonProperty;

public class BadMessage extends SocketMessage
{
	@JsonProperty("reason")
	public String reason;


	public BadMessage(String reason)
	{
		this.reason = reason;
	}


	@Override
	public int getMessageType()
	{
		return SERVER_BAD_MESSAGE;
	}

}
