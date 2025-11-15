'use client';

import { useEffect, useState, useRef } from 'react';

interface Message {
  username: string;
  text: string;
  timestamp: string;
}

export default function Home() {
  const wsRef = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  const [username, setUsername] = useState('');
  const [tempUsername, setTempUsername] = useState('');
  const [text, setText] = useState('');
  const [isConnected, setIsConnected] = useState(false);
  const [showNameModal, setShowNameModal] = useState(true);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };




  
  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  useEffect(() => {
    const websocket = new WebSocket('ws://92.113.34.172:19090/ws');

    websocket.onopen = () => {
      console.log('Conectado ao chat!');
      setIsConnected(true);
    };

    websocket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log('Mensagem recebida:', msg);
      setMessages((prev) => [...prev, msg]);
    };

    websocket.onclose = () => {
      console.log('Desconectado');
      setIsConnected(false);
      wsRef.current = null;
    };

    websocket.onerror = (error) => {
      console.error('Erro no WebSocket:', error);
      setIsConnected(false);
    };

    // store the socket in a ref to avoid calling setState synchronously inside the effect
    wsRef.current = websocket;

    return () => {
      websocket.close();
      wsRef.current = null;
    };
  }, []);

  const sendMessage = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!username.trim()) {
      alert('Por favor, digite seu nome!');
      return;
    }

    if (!text.trim()) return;

    const socket = wsRef.current;
    if (socket && socket.readyState === WebSocket.OPEN) {
      const msg: Message = {
        username: username.trim(),
        text: text.trim(),
        timestamp: new Date().toISOString(),
      };
      console.log('Enviando mensagem:', msg);
      socket.send(JSON.stringify(msg));
      setText('');
    }
  };

  const handleSetUsername = (e: React.FormEvent) => {
    e.preventDefault();
    if (tempUsername.trim()) {
      setUsername(tempUsername.trim());
      setShowNameModal(false);
    }
  };

  const formatTime = (timestamp: string) => {
    const date = new Date(timestamp);
    return date.toLocaleTimeString('pt-BR', { 
      hour: '2-digit', 
      minute: '2-digit' 
    });
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gradient-to-br from-zinc-900 via-zinc-800 to-zinc-900 p-4">
      {/* Modal para definir nome */}
      {showNameModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
          <div className="w-full max-w-md rounded-2xl bg-zinc-900 p-6 shadow-2xl border border-zinc-800">
            <h2 className="mb-4 text-2xl font-bold text-white">Bem-vindo ao Chat! 👋</h2>
            <p className="mb-6 text-sm text-zinc-400">
              Por favor, digite seu nome para começar a conversar
            </p>
            <form onSubmit={handleSetUsername} className="flex flex-col gap-4">
              <input
                type="text"
                value={tempUsername}
                onChange={(e) => setTempUsername(e.target.value)}
                placeholder="Seu nome..."
                autoFocus
                className="rounded-lg bg-zinc-950 px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-blue-500/50 border border-zinc-800"
              />
              <button
                type="submit"
                disabled={!tempUsername.trim()}
                className="rounded-lg bg-gradient-to-r from-blue-500 to-purple-600 px-6 py-3 font-medium text-white hover:from-blue-600 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-blue-500/50 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
              >
                Entrar no Chat
              </button>
            </form>
          </div>
        </div>
      )}

      <div className="flex h-[600px] w-full max-w-4xl flex-col rounded-2xl bg-zinc-950/50 shadow-2xl backdrop-blur-xl border border-zinc-800">
        {/* Header */}
        <div className="flex items-center justify-between border-b border-zinc-800 px-6 py-4">
          <div className="flex items-center gap-3">
            <div className="flex h-10 w-10 items-center justify-center rounded-full bg-gradient-to-br from-blue-500 to-purple-600">
              <span className="text-lg">💬</span>
            </div>
            <div>
              <h1 className="text-lg font-semibold text-white">Chat Global</h1>
              <p className="text-xs text-zinc-400">
                {isConnected ? (
                  <span className="flex items-center gap-1">
                    <span className="h-2 w-2 rounded-full bg-green-500 animate-pulse"></span>
                    Conectado
                  </span>
                ) : (
                  <span className="flex items-center gap-1">
                    <span className="h-2 w-2 rounded-full bg-red-500"></span>
                    Desconectado
                  </span>
                )}
              </p>
            </div>
          </div>
          <div className="text-xs text-zinc-500">
            {messages.length} mensagens
          </div>
        </div>

        {/* Messages Area */}
        <div className="flex-1 overflow-y-auto px-6 py-4 space-y-4">
          {messages.length === 0 ? (
            <div className="flex h-full items-center justify-center">
              <p className="text-zinc-500 text-sm">
                Nenhuma mensagem ainda. Seja o primeiro a enviar! 🚀
              </p>
            </div>
          ) : (
            messages.map((msg, index) => (
              <div
                key={index}
                className="group flex flex-col gap-1 rounded-lg bg-zinc-900/50 p-3 hover:bg-zinc-900/80 transition-colors"
              >
                <div className="flex items-center justify-between">
                  <span className="font-semibold text-sm text-blue-400">
                    {msg.username}
                  </span>
                  <span className="text-xs text-zinc-600 group-hover:text-zinc-500">
                    {formatTime(msg.timestamp)}
                  </span>
                </div>
                <p className="text-zinc-300 text-sm leading-relaxed">
                  {msg.text}
                </p>
              </div>
            ))
          )}
          <div ref={messagesEndRef} />
        </div>

        {/* Input Area */}
        <div className="border-t border-zinc-800 p-4">
          <form onSubmit={sendMessage} className="flex flex-col gap-3">
            <div className="flex items-center gap-2 mb-2">
              <span className="text-sm text-zinc-400">Logado como:</span>
              <span className="text-sm font-semibold text-blue-400">{username}</span>
              <button
                type="button"
                onClick={() => {
                  setShowNameModal(true);
                  setTempUsername('');
                }}
                className="ml-auto text-xs text-zinc-500 hover:text-zinc-300 underline"
              >
                Trocar nome
              </button>
            </div>
            <div className="flex gap-2">
              <input
                type="text"
                value={text}
                onChange={(e) => setText(e.target.value)}
                placeholder="Digite sua mensagem..."
                className="flex-1 rounded-lg bg-zinc-900 px-4 py-3 text-sm text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-blue-500/50 border border-zinc-800"
              />
              <button
                type="submit"
                disabled={!isConnected || !text.trim()}
                className="rounded-lg bg-gradient-to-r from-blue-500 to-purple-600 px-6 py-3 text-sm font-medium text-white hover:from-blue-600 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-blue-500/50 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
              >
                Enviar
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}