'use client';

import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { Mail, Lock } from 'lucide-react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { api } from '@/lib/api';
import type { LoginResponse } from '@/types';

const schema = z.object({
  email: z.string().email('E-mail inválido'),
  password: z.string().min(6, 'A senha deve ter ao menos 6 caracteres'),
});

type FormData = z.infer<typeof schema>;

export default function LoginPage() {
  const router = useRouter();

  const {
    register,
    handleSubmit,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({ resolver: zodResolver(schema) });

  async function onSubmit(data: FormData) {
    try {
      const res = await api.post<LoginResponse>('/api/auth/login', data);
      document.cookie = `auth_token=${res.token}; path=/; max-age=${24 * 60 * 60}; SameSite=Strict`;
      router.push('/');
    } catch {
      setError('root', { message: 'E-mail ou senha inválidos. Tente novamente.' });
    }
  }

  return (
    <div
      className="min-h-screen flex items-center justify-center px-4"
      style={{ background: 'linear-gradient(135deg, #b2dfdb 0%, #e0f2f1 40%, #cde8e5 100%)' }}
    >
      <div className="bg-white rounded-2xl shadow-lg w-full max-w-sm px-8 py-10">
        <div className="flex flex-col items-center mb-6">
          <div className="mb-4">
            <Image
              src="/simbolo-psicologia.jpg"
              alt="Símbolo da Psicologia"
              width={80}
              height={80}
              className="rounded-full object-cover"
              onError={() => {}}
            />
          </div>
          <h1 className="text-2xl font-bold text-slate-800">Gestão Psicológica</h1>
          <p className="text-slate-500 text-sm mt-1">Login de Acesso</p>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-slate-700 mb-1">Email</label>
            <div className="relative">
              <Mail size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
              <input
                type="email"
                placeholder="user@email.com"
                {...register('email')}
                className="w-full pl-9 pr-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent"
              />
            </div>
            {errors.email && <p className="text-red-500 text-xs mt-1">{errors.email.message}</p>}
          </div>

          <div>
            <div className="flex items-center justify-between mb-1">
              <label className="block text-sm font-medium text-slate-700">Senha</label>
              <button type="button" className="text-xs text-teal-600 hover:underline">
                esqueci a senha?
              </button>
            </div>
            <div className="relative">
              <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
              <input
                type="password"
                placeholder="••••••••"
                {...register('password')}
                className="w-full pl-9 pr-4 py-2.5 border border-gray-200 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-teal-500 focus:border-transparent"
              />
            </div>
            {errors.password && <p className="text-red-500 text-xs mt-1">{errors.password.message}</p>}
          </div>

          {errors.root && <p className="text-red-500 text-xs">{errors.root.message}</p>}

          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full bg-teal-600 hover:bg-teal-700 disabled:opacity-60 text-white font-semibold py-3 rounded-lg transition-colors mt-2"
          >
            {isSubmitting ? 'Entrando...' : 'Entrar'}
          </button>
        </form>

        <div className="mt-6 text-center space-y-2">
          <p className="text-sm text-slate-500">
            <button className="hover:underline">Esqueceu a senha?</button>
          </p>
          <p className="text-sm text-slate-500">
            <button className="hover:underline">Criar Nova Conta</button>
          </p>
        </div>
      </div>
    </div>
  );
}
