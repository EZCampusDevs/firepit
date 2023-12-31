package org.ezcampus.firepit.data.message;

import org.ezcampus.firepit.data.Client;

import com.fasterxml.jackson.annotation.JsonProperty;

public class RenameMessage extends SocketMessage
{
	@JsonProperty("new_name")
	public String newName;
	
	@JsonProperty("client")
	public Client client;

	public RenameMessage(Client c, String newname)
	{
		this.client = c;
		this.newName = newname;
	}

	@Override
	public int getMessageType()
	{
		return SET_CLIENT_NAME;
	}

}
